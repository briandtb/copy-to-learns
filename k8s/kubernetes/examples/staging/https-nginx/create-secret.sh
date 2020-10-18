#!/bin/bash
cat <<EOF > $3
apiVersion: v1
kind: Secret
metadata:
  name: nginxsecret
  namespace: default
type: kubernetes.io/tls
data:
  tls.key: $(cat $1 | base64 | tr -d '\n')
  tls.crt: $(cat $2 | base64 | tr -d '\n')
EOF 
