terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }

    local = {
      source  = "hashicorp/local"
      version = "2.1.0"
    }

    null = {
      source  = "hashicorp/null"
      version = "3.1.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.1"
    }
  }

  required_version = ">= 0.14.9"
  backend "s3" {
    bucket = "unity-demo-state"
    key    = "demo/state"
    region = "us-east-1"
    profile = "saml-pub"
  }
}

provider "aws" {
  profile = "saml-pub"
}

resource "aws_elasticsearch_domain" "unity-sample" {
  domain_name           = "unityexample"
  elasticsearch_version = 7.11
  cluster_config {
    instance_type = "i2.xlarge.elasticsearch"
    instance_count = 2
    zone_awareness_enabled = true
  }
  tags = {
    hello = "world", someothertag="no"
  }
}