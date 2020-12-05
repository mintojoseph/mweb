resource "helm_release" "minikube" {
  name          = "mweb"
  chart         = "../mweb-helm"
}