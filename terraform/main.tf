resource "azurerm_resource_group" "group" {
  name     = var.installation
  location = var.location
}

resource "azurerm_public_ip" "ip" {
  name                = var.installation
  resource_group_name = azurerm_resource_group.group.name
  location            = azurerm_resource_group.group.location
  allocation_method   = "Static"
}

resource "azurerm_kubernetes_cluster" "aks" {
  name                = var.installation
  location            = azurerm_resource_group.group.location
  resource_group_name = azurerm_resource_group.group.name
  dns_prefix          = "tabbycats-demo"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }
}
