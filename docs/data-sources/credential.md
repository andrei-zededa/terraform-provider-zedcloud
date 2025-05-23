---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zedcloud_credential Data Source - terraform-provider-zedcloud"
subcategory: ""
description: |-
  
---

# zedcloud_credential (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `list` (Block List) (see [below for nested schema](#nestedblock--list))
- `next` (Block List) (see [below for nested schema](#nestedblock--next))
- `summary_by_state` (Block List) (see [below for nested schema](#nestedblock--summary_by_state))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--list"></a>
### Nested Schema for `list`

Optional:

- `current_cred` (String, Sensitive)
- `forgot` (Boolean)
- `new_cred` (String, Sensitive)
- `owner` (String)
- `salt` (Number)
- `type` (String)

Read-Only:

- `id` (String)


<a id="nestedblock--next"></a>
### Nested Schema for `next`

Optional:

- `order_by` (List of String) OrderBy helps in sorting the list response
- `page_num` (Number) Page Number
- `page_size` (Number) Defines the page size
- `page_token` (String) Page Token
- `total_pages` (Number) Total number of pages to be fetched.


<a id="nestedblock--summary_by_state"></a>
### Nested Schema for `summary_by_state`

Optional:

- `description` (String) Summary description
- `total` (Number) Total
- `values` (Map of Number) Values: Map for storing <string, uint32>
