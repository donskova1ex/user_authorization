// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * user_authorization service - OpenAPI 3.0
 *
 * Сервис авторизации пользователей
 *
 * API version: 1.0.0
 */

package openapi

type Role struct {
	Uuid string `json:"uuid,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

// AssertRoleRequired checks if the required fields are not zero-ed
func AssertRoleRequired(obj Role) error {
	return nil
}

// AssertRoleConstraints checks if the values respects the defined constraints
func AssertRoleConstraints(obj Role) error {
	return nil
}