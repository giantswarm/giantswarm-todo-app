{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "giantswarm-todo.chart-name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "giantswarm-todo.chart-fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "giantswarm-todo.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "giantswarm-todo.labels" -}}
{{- if .Name -}}
app.kubernetes.io/name: {{ .Name }}
{{ end -}}
helm.sh/chart: {{ include "giantswarm-todo.chart" .Root }}
app.kubernetes.io/instance: {{ .Root.Release.Name }}
{{- if .Root.Chart.AppVersion }}
app.kubernetes.io/version: {{ .Root.Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Root.Release.Service }}
{{- end -}}

{{/*
Antiaffinity for pods
*/}}
{{- define "giantswarm-todo.antiaffinity" -}}
podAntiAffinity:
  preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 50
      podAffinityTerm:
        topologyKey: "kubernetes.io/hostname"
        labelSelector:
          matchExpressions:
            - key: "app.kubernetes.io/name"
              operator: In
              values:
                - {{ .Name | quote}}
            - key: "app.kubernetes.io/instance"
              operator: In
              values:
                - {{ .Root.Release.Name }}
    - weight: 100
      podAffinityTerm:
        topologyKey: "failure-domain.beta.kubernetes.io/zone"
        labelSelector:
          matchExpressions:
            - key: "app.kubernetes.io/name"
              operator: In
              values:
                - {{ .Name | quote}}
            - key: "app.kubernetes.io/instance"
              operator: In
              values:
                - {{ .Root.Release.Name }}
{{- end -}}

{{/*
Generic labels
*/}}
{{- define "giantswarm-todo.generic.labels" -}}
{{ $data := dict "Root" $ "Name" "" }}
{{ include "giantswarm-todo.labels" $data }}
{{- end -}}

{{/*
Todomanager labels
*/}}
{{- define "giantswarm-todo.todomanager.labels" -}}
{{ $data := dict "Root" $ "Name" "todomanager" }}
{{ include "giantswarm-todo.labels" $data }}
{{- end -}}

{{/*
Todomanager match labels
*/}}
{{- define "giantswarm-todo.todomanager.match-labels" -}}
app.kubernetes.io/name: todomanager
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Todomanager antiaffinity
*/}}
{{- define "giantswarm-todo.todomanager.antiaffinity" -}}
{{ $data := dict "Root" $ "Name" "todomanager" }}
{{ include "giantswarm-todo.antiaffinity" $data }}
{{- end -}}

{{/*
Apiserver labels
*/}}
{{- define "giantswarm-todo.apiserver.labels" -}}
{{ $data := dict "Root" $ "Name" "apiserver" }}
{{ include "giantswarm-todo.labels" $data }}
{{- end -}}

{{/*
Apiserver match labels
*/}}
{{- define "giantswarm-todo.apiserver.match-labels" -}}
app.kubernetes.io/name: apiserver
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Apiserver antiaffinity
*/}}
{{- define "giantswarm-todo.apiserver.antiaffinity" -}}
{{ $data := dict "Root" $ "Name" "apiserver" }}
{{ include "giantswarm-todo.antiaffinity" $data }}
{{- end -}}