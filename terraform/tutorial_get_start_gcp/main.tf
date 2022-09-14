terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file("~/sandbox-janwonkim-5584d199acb9.json")

  project = "sandbox-janwonkim"
  region  = "us-central"
  zone    = "us-central"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}
