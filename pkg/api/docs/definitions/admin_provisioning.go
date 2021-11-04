package apidocs

// swagger:route POST /admin/provisioning/dashboards/reload admin_provisioning reloadProvisionedDashboards
//
// Reload dashboard provisioning configurations.
//
// Reloads the provisioning config files for dashboards again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.
// If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:dashboards`.
//
// Security:
// - basic:
//
// Responses:
// 200: okResponse
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError

// swagger:route POST /admin/provisioning/datasources/reload admin_provisioning reloadProvisionedDatasources
//
// Reload datasource provisioning configurations.
//
// Reloads the provisioning config files for datasources again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.
// If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:datasources`.
//
// Security:
// - basic:
//
// Responses:
// 200: okResponse
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError

// swagger:route POST /admin/provisioning/plugins/reload admin_provisioning reloadProvisionedPlugins
//
// Reload plugin provisioning configurations.
//
// Reloads the provisioning config files for plugins again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.
// If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:plugin`.
//
// Security:
// - basic:
//
// Responses:
// 200: okResponse
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError

// swagger:route POST /admin/provisioning/notifications/reload admin_provisioning reloadProvisionedAlertNotifiers
//
// Reload legacy alert notifier provisioning configurations.
//
// Reloads the provisioning config files for legacy alert notifiers again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.
// If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:notifications`.
//
// Security:
// - basic:
//
// Responses:
// 200: okResponse
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError

// swagger:route POST /admin/provisioning/accesscontrol/reload admin_provisioning reloadProvisionedAccessControl
//
// Reload access control provisioning configurations.
//
// Reloads the provisioning config files for access control again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning.
// If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:accesscontrol`.
//
// Security:
// - basic:
//
// Responses:
// 200: okResponse
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError
