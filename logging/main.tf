resource "kubernetes_namespace" "grafana" {
  metadata {
    annotations = {
      name = "grafana"
    }
  name = "grafana"
  }
}

resource "helm_release" "grafana" {
  name       = "grafana"
  repository = "https://grafana.github.io/helm-charts"
  chart      = "grafana"
  namespace  = "${kubernetes_namespace.grafana.metadata.0.name}"
  force_update = true

  set {
    name  = "service.type"
    value = "LoadBalancer"
  }
}

resource "helm_release" "loki" {
  name       = "loki"
  repository = "https://grafana.github.io/loki/charts"
  chart      = "loki-stack"
  force_update = true
}