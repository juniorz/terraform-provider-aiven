package aiven

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceCassandra() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceServiceRead,
		Schema:      resourceSchemaAsDatasourceSchema(cassandraSchema(), "project", "service_name"),
	}
}
