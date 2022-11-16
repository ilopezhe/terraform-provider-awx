AWX Terraform Provider
======================

An autogenerated terraform provider based on the API specifications as provided by the `/api/v2/` endpoint.

TODO:
-----
* Unit tests
* Integration tests

Implementation status
---------------------

* [x] Inventory
  * [x] Object roles
* [x] Project
  * [x] Object roles
* [x] Organization
  * [x] Object roles
  * [x] Associate/disassociate with Execution Environment
  * [x] Associate/disassociate with Instance Group
* [x] Team
  * [x] Object roles
  * [x] Associate/disassociate roles
* [x] ExecutionEnvironment
* [x] CredentialType
* [x] User
  * [x] Associate/disassociate roles
* [x] Credential input sources
* [x] Credential
  * [x] Object roles
* [x] Applications
* [x] Labels (not possible to delete, deleting the organization deletes the labels for the organization)
* [x] Instance Groups
* [x] Schedules
* [x] Groups
* [x] Tokens
* [x] Me
* [x] Host
  * [x] Associate/disassociate with Group
* [x] Ad Hoc Commands
* [x] Notification Templates
* [x] Settings
  * [x] System - Define system-level features and functions
    * [x] Miscellaneous System settings
    * [x] Miscellaneous Authentication settings
    * [x] Logging settings
  * [x] User Interface - Set preferences for data collection, logos, and logins
    * [x] User Interface settings
  * [x] Jobs - Update settings pertaining to Jobs within AWX
    * [x] Jobs settings
  * [x] Authentication - Enable simplified login for your AWX applications
    * [x] Azure AD settings
    * [x] GitHub settings
    * [x] Google OAuth 2 settings
    * [x] LDAP settings
    * [x] SAML settings
    * [x] Generic OIDC settings
* [x] Inventory sources
* [x] Job Template
  * [x] Object roles
  * [x] Survey Spec
  * [x] Notifications
  * [x] Associate/disassociate with Credential
* [x] Workflow Job Template
  * [x] Object roles
  * [x] Survey Spec
  * [ ] Nodes
  * [x] Notifications
