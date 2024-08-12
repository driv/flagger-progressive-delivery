This repo is an example for the [blog post]() //TODO Add link

### Requirements
- docker
- kind
- kubectl
- [cloud-provider-kind](https://github.com/kubernetes-sigs/cloud-provider-kind?tab=readme-ov-file#install)

# Cluster
You can use any cluster you have available
```
kind create cluster --name=kind-progressive-delivery
```

# FluxCD bootstrap
```
kubectl apply -k https://github.com/driv/flagger-progressive-delivery/clusters/my-cluster/infrastructure/flux-system
```

# External IP.

The `nginx-ingress` installation creates a `LoadBalancer`. For this `LoadBalancer` to be accessible, we need to assign it an external IP. This is where `cloud-provider-kind` comes in.
```
# Run the command in a separate terminal and leave it running
./cloud-provider-kind
```

Once `cloud-provider-kind` is running, we can configure our `Canary` and `Ingress`.

```
EXTERNAL_IP=$(kubectl get service -n ingress-nginx ingress-nginx-controller -o jsonpath='{ .status.loadBalancer.ingress[0].ip }' | sed 's/\./-/g' ) &&
echo $EXTERNAL_IP &&
kubectl -n default get canary golang-api -o yaml | sed "s/172-18-0-3/$EXTERNAL_IP/g" | kubectl replace -f - &&
kubectl -n default get ingress number-generator -o yaml | sed "s/172-18-0-3/$EXTERNAL_IP/g" | kubectl replace -f - ;
```
# Frontend
```
echo "Visit https://www.$EXTERNAL_IP.nip.io/"
```
Click on the link and access the frontend (The browser will complain about the ssl certificate).

# Update the backend

Now it's time to play with Flagger.

There are 4 image tags available.

- `next`: It generates n+1 number, it's the one currently running.
- `sleep`: Same as next but it'll sleep for 2 seconds before returning the result. It'll fail the release.
- `good_fibonacci`: It returns the nth number from the series. With numbers up to 100 it can keep the response time within the limits of the [analysis](https://github.com/driv/flagger-progressive-delivery/blob/main/clusters/my-cluster/apps/golang-api/canary.yaml#L33).
- `bad_fibonacci`: It returns the nth number from the series. It's an extremely inefficient implementation. It'll fail the release, but only if "real" traffic from the UI is coming in, the [load test](https://github.com/driv/flagger-progressive-delivery/blob/main/clusters/my-cluster/apps/golang-api/canary.yaml#L55) will not detect how slow it is with big numbers. (keep the UI generating requests to see the release fail)

Update the deployment image:
```
kubectl -n default set image deployment/golang-api golang-api=driv/buildpack-playground-golang-api:<image_tag>
```

Check the release progress in the Flagger logs:
```
kubectl logs -n ingress-nginx --selector app.kubernetes.io/instance=flagger -f
```
Or the Kubernetes events:
```
kubectl -n default get events --watch
```