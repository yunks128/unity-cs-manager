provider "aws" {
  profile = "saml-pub"
}

resource "aws_elasticsearch_domain" "unity-sample" {
  domain_name           = "unityexample"
}