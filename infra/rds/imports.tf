import {
  id = "asterisk-db"
  to = module.asterisk-db.aws_db_instance.this
}

import {
  id = "infra-db"
  to = module.infra-db.aws_db_instance.this
}

import {
  id = "intranet-dev-db"
  to = module.intranet-dev-db.aws_db_instance.this
}

import {
  id = "intra-prod-restored"
  to = module.intra-prod-restored.aws_db_instance.this
}

import {
  id = "prod-intranet-db"
  to = module.prod-intranet-db.aws_db_instance.this
}

import {
  id = "read-replica"
  to = module.read-replica.aws_db_instance.this
}

