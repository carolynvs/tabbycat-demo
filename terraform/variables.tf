variable "installation" {
  description = "Name of the bundle installation"
  type        = string
}

variable "location" {
  description = "Azure location"
  type        = string
}

variable "failover_location" {
  description = "Azure location"
  type        = string
  default     = "eastus2"
}
