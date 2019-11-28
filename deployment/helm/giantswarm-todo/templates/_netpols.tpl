{{/*
Network policy: kube-dns
*/}}
{{- define "giantswarm-todo.netpol.kube-dns" -}}
- namespaceSelector:
    matchLabels:
      name: kube-system
- podSelector:
    matchLabels:
      k8s-app: kube-dns
ports:
- protocol: TCP
  port: 53
- protocol: UDP
  port: 53
{{- end -}}

{{/*
Network policy: linkerd
*/}}
{{- define "giantswarm-todo.netpol.linkerd" -}}
- namespaceSelector:
    matchLabels:
      linkerd.io/control-plane-ns: {{ .Values.linkerdNamespace }}
{{- end -}}