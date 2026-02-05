module "asterisk-db" {
  source            = "./modules/rds-padrao"
  allocated_storage = 10
  engine            = "postgres"
  engine_version    = "16.8"
  identifier        = "asterisk-db"
  instance_class    = "db.t4g.small"
  multi_az          = false
  storage_type      = "gp2"
  environment       = "production"
  terraform_managed = true
}

module "infra-db" {
  source            = "./modules/rds-padrao"
  allocated_storage = 100
  engine            = "postgres"
  engine_version    = "16.8"
  identifier        = "infra-db"
  instance_class    = "db.t4g.medium"
  multi_az          = true
  storage_type      = "gp3"
  environment       = "production"
  terraform_managed = true
}

module "intranet-dev-db" {
  source            = "./modules/rds-padrao"
  allocated_storage = 150
  engine            = "postgres"
  engine_version    = "16.8"
  identifier        = "intranet-dev-db"
  instance_class    = "db.t4g.micro"
  multi_az          = false
  storage_type      = "gp2"
  environment       = "production"
  terraform_managed = true
}

module "intra-prod-restored" {
  source            = "./modules/rds-padrao"
  allocated_storage = 200
  engine            = "postgres"
  engine_version    = "13.20"
  identifier        = "intra-prod-restored"
  instance_class    = "db.m6g.2xlarge"
  multi_az          = true
  storage_type      = "gp3"
  environment       = "production"
  terraform_managed = true
}

module "prod-intranet-db" {
  source            = "./modules/rds-padrao"
  allocated_storage = 200
  engine            = "postgres"
  engine_version    = "16.8"
  identifier        = "prod-intranet-db"
  instance_class    = "db.m6g.large"
  multi_az          = true
  storage_type      = "gp3"
  environment       = "production"
  terraform_managed = true
}

module "read-replica" {
  source            = "./modules/rds-padrao"
  allocated_storage = 200
  engine            = "postgres"
  engine_version    = "13.20"
  identifier        = "read-replica"
  instance_class    = "db.t4g.small"
  multi_az          = false
  storage_type      = "gp3"
  environment       = "production"
  terraform_managed = true
}

