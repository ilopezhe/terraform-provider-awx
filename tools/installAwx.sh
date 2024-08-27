#!/bin/bash
echo "==> Install Operator..."
# kubectl apply -f https://raw.githubusercontent.com/ansible/awx-operator/devel/deploy/awx-operator.yaml

curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3

chmod +x get_helm.sh

./get_helm.sh

helm version

helm install --create-namespace --namespace ansible-awx awx-operator "https://github.com/ansible/awx-operator/releases/download/2.19.1/awx-operator-2.19.1.tgz"

echo "==> Wait Operator started..."
kubectl wait --for=condition=ready pod -l control-plane=controller-manager -n ansible-awx  --timeout=120s

echo "==> Starting AWX Test installation..."
# kubectl create ns ansible-awx
# kubectl delete AWX awx -n ansible-awx
KUBENODE_IP=$(kubectl get node -ojson | jq '.items[0].status.addresses[0].address' -r)
INGRESS_DOMAIN="awx.localhost.local"

kubectl create secret generic admin-credentials \
  --from-literal=password=changeme \
  --namespace=ansible-awx

cat <<EOF | kubectl apply -f -
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx
  namespace: ansible-awx
spec:
  deployment_type: awx
  admin_user: test
  admin_email: test@example.com
  admin_password_secret: admin-credentials
  ingress_type: Ingress
  hostname: ${INGRESS_DOMAIN}
EOF
# broadcast_websocket_secret: admin-credentials

echo "==> Wait for the Operator to start the AWX Deployment..."
sleep 60

echo "==> Wait for AWX deployment to start..."
kubectl wait --for=condition=ready pod -l "app.kubernetes.io/name=awx-web" -n ansible-awx --timeout=240s

echo "==> Wait for AWX task to finish migrations..."
kubectl wait --for=condition=ready pod -l "app.kubernetes.io/name=awx-task" -n ansible-awx --timeout=300s

echo "==> Run make port-forward to start a port-forward from 8080 to 80 on the service"
