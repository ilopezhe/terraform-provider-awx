resource "awx_organization" "this" {
  name = "acc-organization"
}

resource "awx_inventory" "this" {
  name            = "acc-inventory"
  organization = awx_organization.this.id
  variables       = jsonencode({
    system_supporters = ["pi"]
  })
}

# TODO FIX
resource "awx_inventory" "smart" {
  name            = "acc-smart-inventory"
  organization = awx_organization.this.id
  host_filter = "name__icontains=localhost"
  kind = "smart"
  variables       = jsonencode({
    system_supporters = ["pi"]
  })
}

data "awx_credential_type" "machine" {
  id = 1
}

resource "awx_credential" "credential" {
  name            = "acc-machine-credential"
  credential_type = data.awx_credential_type.machine.id
  organization = awx_organization.this.id

  inputs = jsonencode({
    username            = "pi"
    ssh_key_data = file("${path.module}/files/id_rsa")
    ssh_public_key_data = file("${path.module}/files/id_rsa.pub")
    ssh_key_unlock      = "test"
  })
}

resource "awx_workflow_job_template" "this" {
  name            = "acc-workflow-job"
  organization = awx_organization.this.id
  inventory    = awx_inventory.this.id
  ask_inventory_on_launch = true
}
# resource "random_uuid" "workflow_node_k3s_uuid" {}

# resource "awx_workflow_job_template_node" "this" {

#   workflow_job_template_id = awx_workflow_job_template.this.id
#   unified_job_template  = awx_job_template.this.id
#   inventory             = awx_inventory.this.id
#   identifier               = random_uuid.workflow_node_k3s_uuid.result
# }
# resource "random_uuid" "workflow_node_second_uuid" {}

# resource "awx_workflow_job_template_node_success" "this" {

#   workflow_job_template_node_id = awx_workflow_job_template_node.this.id
#   workflow_job_template_id       = awx_workflow_job_template.this.id
#   unified_job_template  = awx_job_template.this.id
#   inventory                  = awx_inventory.this.id
#   identifier                    = random_uuid.workflow_node_second_uuid.result
# }

resource "awx_host" "this" {
  name         = "acc-node1"
  inventory = awx_inventory.this.id
  enabled      = true
  variables = jsonencode({
    ansible_host = "192.168.178.29"
  })
}

resource "awx_project" "this" {
  name                 = "acc-project"
  scm_type             = "git"
  scm_url              = "https://github.com/ansible/ansible-tower-samples"
  scm_update_on_launch = true
  organization      = awx_organization.this.id
}

## give Certsmanger Time to Work
resource "time_sleep" "wait_seconds" {
  depends_on = [awx_project.this]

  create_duration = "30s"
}

resource "awx_job_template" "this" {
  depends_on     = [time_sleep.wait_seconds]
  name           = "acc-job-template"
  job_type       = "run"
  inventory   = 1
  project     = awx_project.this.id
  playbook       = "hello_world.yml"
  become_enabled = true
  ask_inventory_on_launch = true
}

# resource "awx_job_template_credential" "this" {
#   job_template_id = awx_job_template.this.id
#   credential_id   = awx_credential_machine.credential.id
# }

resource "awx_inventory_source" "this" {
  name              = "acc-inventory-sources"
  inventory      = awx_inventory.this.id
  source_project = awx_project.this.id
  source = "scm"
}


output "inventory" {
  value = awx_inventory.this.id
}

# data "awx_project_role" "this" {
#   name       = "Admin"
#   project = awx_project.this.id
# }

# data "awx_inventory_role" "this" {
#   name       = "Read"
#   inventory = awx_inventory.this.id
# }

# data "awx_job_template_role" "this" {
#   name       = "Read"
#   job_template_id = awx_job_template.this.id
# }

resource "awx_team" "this" {
  organization = 1
  name = "acc-team"

  # role_entitlement {
  #   role_id = data.awx_project_role.this.id
  # }

  # role_entitlement {
  #   role_id = data.awx_inventory_role.this.id
  # }

  # role_entitlement {
  #   role_id = data.awx_job_template_role.this.id
  # }
}

resource "awx_credential" "aws" {
  name               = "acc-credential"
  credential_type = 5
  organization    = 1

  inputs = jsonencode({
    username = "something"
    password = "other"
  })
}

resource "awx_credential_type" "hashicorp_vault" {
  name        = "acc-credential-type"
  kind        = "cloud"

  inputs = jsonencode({
    fields = [
      {
        "type"  = "string",
        "id"    = "vault_addr",
        "label" = "Vault url"
      },
      {
        "type"   = "string",
        "id"     = "approle_secret_id",
        "label"  = "Vault secret",
        "secret" = true
      },
      {
        "type"  = "string",
        "id"    = "mount_point",
        "label" = "Vault mount point"
      },
      {
        "type"  = "string",
        "id"    = "approle_id",
        "label" = "Vault role id"
      }
    ]
    required = [
      "vault_addr",
      "approle_secret_id",
      "mount_point",
      "approle_id"
    ]
  })

  injectors = jsonencode({
    "env" = {
      "VAULT_ADDR"              = "{{ vault_addr }}",
      "VAULT_MOUNT_POINT"       = "{{ mount_point }}",
      "VAULT_APPROLE_ID"        = "{{ approle_id }}",
      "VAULT_APPROLE_SECRET_ID" = "{{ approle_secret_id }}"
    }
  })
}

resource "awx_schedule" "this" {
  name                    = "acc-schedule"
  rrule                   = "DTSTART;TZID=Europe/Paris:20501214T120000 RRULE:INTERVAL=1;FREQ=DAILY"
  unified_job_template = awx_job_template.this.id
  inventory               = awx_inventory.this.id
}


data "awx_project_object_roles" "this" {
  id = awx_project.this.id
}

output "awx_project_object_roles" {
  value = length(data.awx_project_object_roles.this.roles)
}

data "awx_job_template_object_roles" "this" {
  id = awx_job_template.this.id
}

output "awx_job_template_object_roles" {
  value = length(data.awx_job_template_object_roles.this.roles)
}

data "awx_inventory_object_roles" "this" {
  id = awx_inventory.this.id
}

output "awx_inventory_object_roles" {
  value = length(data.awx_inventory_object_roles.this.roles)
}






# data "awx_credential" "this" {
#   name = "acc-credential"
# }

# output "awx_credential" {
#   value = data.awx_credential.this.id
# }

# data "awx_credential_type" "this" {
#   name = "acc-credential-type"
# }

# output "awx_credential_type" {
#   value = data.awx_credential_type.this.id
# }

data "awx_inventory" "this" {
  name = awx_inventory.this.name
}

output "awx_inventory" {
  value = data.awx_inventory.this.id
}

# data "awx_inventory_group" "this" {
#   name = "acc-inventory-group"
# }

# output "awx_inventory_group" {
#   value = data.awx_inventory_group.this.id
# }

data "awx_job_template" "this" {
  name = awx_job_template.this.name
}

output "awx_job_template" {
  value = data.awx_job_template.this.id
}

data "awx_organization" "this" {
  name = awx_organization.this.name
}

output "awx_organization" {
  value = data.awx_organization.this.id
}

data "awx_project" "this" {
  id = awx_project.this.id
}

output "awx_project" {
  value = data.awx_project.this.id
}

data "awx_schedule" "this" {
  id = awx_schedule.this.id
}

output "awx_schedule" {
  value = data.awx_schedule.this.id
}

data "awx_team" "this" {
  name = awx_team.this.name
}

output "awx_team" {
  value = data.awx_team.this.id
}
