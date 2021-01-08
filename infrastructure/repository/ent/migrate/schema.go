// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// SkuActionColumns holds the columns for the "sku_action" table.
	SkuActionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "app_id", Type: field.TypeString},
		{Name: "sku_name", Type: field.TypeString},
		{Name: "sku_name_cn", Type: field.TypeString},
		{Name: "action", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// SkuActionTable holds the schema information for the "sku_action" table.
	SkuActionTable = &schema.Table{
		Name:        "sku_action",
		Columns:     SkuActionColumns,
		PrimaryKey:  []*schema.Column{SkuActionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SkuConfigColumns holds the columns for the "sku_config" table.
	SkuConfigColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "config", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "sku_action_id", Type: field.TypeInt, Nullable: true},
	}
	// SkuConfigTable holds the schema information for the "sku_config" table.
	SkuConfigTable = &schema.Table{
		Name:       "sku_config",
		Columns:    SkuConfigColumns,
		PrimaryKey: []*schema.Column{SkuConfigColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "sku_config_sku_action_config",
				Columns: []*schema.Column{SkuConfigColumns[5]},

				RefColumns: []*schema.Column{SkuActionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkuStatusCodeColumns holds the columns for the "sku_status_code" table.
	SkuStatusCodeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "insurance_code", Type: field.TypeString},
		{Name: "reflect_code", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "sku_action_id", Type: field.TypeInt, Nullable: true},
	}
	// SkuStatusCodeTable holds the schema information for the "sku_status_code" table.
	SkuStatusCodeTable = &schema.Table{
		Name:       "sku_status_code",
		Columns:    SkuStatusCodeColumns,
		PrimaryKey: []*schema.Column{SkuStatusCodeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "sku_status_code_sku_action_status_code",
				Columns: []*schema.Column{SkuStatusCodeColumns[6]},

				RefColumns: []*schema.Column{SkuActionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkuTemplateColumns holds the columns for the "sku_template" table.
	SkuTemplateColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "request", Type: field.TypeString},
		{Name: "response", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "sku_action_id", Type: field.TypeInt, Nullable: true},
	}
	// SkuTemplateTable holds the schema information for the "sku_template" table.
	SkuTemplateTable = &schema.Table{
		Name:       "sku_template",
		Columns:    SkuTemplateColumns,
		PrimaryKey: []*schema.Column{SkuTemplateColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "sku_template_sku_action_template",
				Columns: []*schema.Column{SkuTemplateColumns[6]},

				RefColumns: []*schema.Column{SkuActionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SkuActionTable,
		SkuConfigTable,
		SkuStatusCodeTable,
		SkuTemplateTable,
	}
)

func init() {
	SkuConfigTable.ForeignKeys[0].RefTable = SkuActionTable
	SkuStatusCodeTable.ForeignKeys[0].RefTable = SkuActionTable
	SkuTemplateTable.ForeignKeys[0].RefTable = SkuActionTable
}
