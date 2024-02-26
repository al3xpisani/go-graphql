```GRAPHQL PLAYGROUND
mutation createCategory {
  createCategory(
    input: {name: "Electronics", description: "electronics level a,b,c"}
  ) {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(
    input: {name: "Curso matemática", description: "Curso muito bom", categoryId: 1}
  ) {
    id
    name
    description
  }
}

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    description
    courses {
      name
    }
  }
}

query queryCourses {
  courses {
    id
    name
    description
  }
}

query queryCoursesWithCategories {
  courses {
    id
    name
    description
    category {
      name
      description
    }
  }
}
```
