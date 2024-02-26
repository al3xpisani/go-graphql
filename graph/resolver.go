//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/al3xpisani/GO-GRAPHQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
