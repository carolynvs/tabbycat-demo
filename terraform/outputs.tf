output "static_ip" {
  value = azurerm_public_ip.ip.ip_address
}

output "kubeconfig" {
  value = azurerm_kubernetes_cluster.aks.kube_config_raw
  sensitive = true
}

output "connstr" {
  value = azurerm_cosmosdb_account.db.connection_strings
  sensitive = true
}
