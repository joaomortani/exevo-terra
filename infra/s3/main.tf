module "aws-athena-query-results-394188845576-us-east-1" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "aws-athena-query-results-394188845576-us-east-1"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "aws-cloudtrail-logs-394188845576-c16a75e2" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "aws-cloudtrail-logs-394188845576-c16a75e2"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "aws-glue-assets-394188845576-us-east-1" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "aws-glue-assets-394188845576-us-east-1"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "aws-load-balancer-logs-394188845576" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "aws-load-balancer-logs-394188845576"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "cf-templates-1rpthmgggouc5-us-east-1" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "cf-templates-1rpthmgggouc5-us-east-1"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "content-service-prod-access-logs-bucket" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "content-service-prod-access-logs-bucket"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "copy-teste-origem" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "copy-teste-origem"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "corujaintra-394188845576-athena-bucket" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "corujaintra-394188845576-athena-bucket"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "dev-estrategia-intranet-files" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "dev-estrategia-intranet-files"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "estrategia-intra-terraform-bucket" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "estrategia-intra-terraform-bucket"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "estrategiatranscript" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "estrategiatranscript"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "intra-odata-log-stream" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "intra-odata-log-stream"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "intranet-datasets" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "intranet-datasets"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "intranet-frontend-imagens" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "intranet-frontend-imagens"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "intranet.estrategia-sandbox.com.br" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "intranet.estrategia-sandbox.com.br"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "intranet.estrategia.com" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "intranet.estrategia.com"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "prod-estrategia-intranet-files" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "prod-estrategia-intranet-files"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "prod-intranet-stream" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "prod-intranet-stream"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "protego-fsp-394188845576" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "protego-fsp-394188845576"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "retranscode" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "retranscode"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "testes-intranet-estrategia" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "testes-intranet-estrategia"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

module "tf-trail-prod-intra" {
  source                   = "terraform-aws-modules/s3-bucket/aws"
  bucket                   = "tf-trail-prod-intra"
  acl                      = "private"
  control_object_ownership = true
  object_ownership         = "ObjectWriter"
  versioning               = "map[enabled:true]"
}

