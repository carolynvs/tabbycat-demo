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

resource "random_integer" "rand" {
  min = 10000
  max = 99999
}

resource "azurerm_cosmosdb_account" "db" {
  name                = "${var.installation}-${random_integer.rand.result}"
  resource_group_name = azurerm_resource_group.group.name
  location            = var.location
  offer_type          = "Standard"
  kind                = "MongoDB"

  capabilities {
    name = "EnableMongo"
  }

  geo_location {
    location          = azurerm_resource_group.group.location
    failover_priority = 0
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 10
    max_staleness_prefix    = 200
  }
}

resource "azurerm_cosmosdb_mongo_database" "db" {
  name                = var.installation
  resource_group_name = azurerm_resource_group.group.name
  account_name        = azurerm_cosmosdb_account.db.name
}

resource "azurerm_cosmosdb_mongo_collection" "db_bookmarks" {
  name                = "bookmarks"
  resource_group_name = azurerm_resource_group.group.name
  account_name        = azurerm_cosmosdb_account.db.name
  database_name       = azurerm_cosmosdb_mongo_database.db.name
}
