#!/usr/bin/env bash
echo "Installing security profiles operator"

echo "security operator is installed, mainly for seccomp, in the future will be also for app armor"

echo "Reference: https://github.com/kubernetes-sigs/security-profiles-operator/blob/main/installation-usage.md"

echo "Cert manager"

kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.7.1/cert-manager.yaml

echo "wait for running cert manager"
kubectl --namespace cert-manager wait --for condition=ready pod -l app.kubernetes.io/instance=cert-manager

echo "operator"
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/security-profiles-operator/main/deploy/operator.yaml

echo "in case of production apply profile using the daemon and profile with the corresponding namespace"