// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * user_authorization service - OpenAPI 3.0
 *
 * Сервис авторизации пользователей
 *
 * API version: 1.0.0
 */

package openapi

type Permission struct {
	Uuid string `json:"uuid,omitempty"`

	PermissionsName string `json:"permissions_name"`
}

// AssertPermissionRequired checks if the required fields are not zero-ed
func AssertPermissionRequired(obj Permission) error {
	elements := map[string]interface{}{
		"permissions_name": obj.PermissionsName,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPermissionConstraints checks if the values respects the defined constraints
func AssertPermissionConstraints(obj Permission) error {
	return nil
}