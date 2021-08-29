output "resource_group" {
  value = azurerm_resource_group.group.name
}

output "kubeconfig" {
  value = azurerm_kubernetes_cluster.aks.kube_config_raw
  sensitive = true
}

output "connstr" {
  value = azurerm_cosmosdb_account.db.connection_strings[0]
  sensitive = true
}
