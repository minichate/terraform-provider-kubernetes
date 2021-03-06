// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	redis "google.golang.org/api/redis/v1beta1"
)

func resourceRedisInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisInstanceCreate,
		Read:   resourceRedisInstanceRead,
		Update: resourceRedisInstanceUpdate,
		Delete: resourceRedisInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(360 * time.Second),
			Update: schema.DefaultTimeout(360 * time.Second),
			Delete: schema.DefaultTimeout(360 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"memory_size_gb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alternative_location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"authorized_network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"redis_configs": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"redis_version": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"reserved_ip_range": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"tier": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BASIC", "STANDARD_HA", ""}, false),
				Default:      "BASIC",
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_location_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRedisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alternative_location_id"); !isEmptyValue(reflect.ValueOf(alternativeLocationIdProp)) && (ok || !reflect.DeepEqual(v, alternativeLocationIdProp)) {
		obj["alternativeLocationId"] = alternativeLocationIdProp
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(memorySizeGbProp)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_version"); !isEmptyValue(reflect.ValueOf(redisVersionProp)) && (ok || !reflect.DeepEqual(v, redisVersionProp)) {
		obj["redisVersion"] = redisVersionProp
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1/projects/{{project}}/locations/{{region}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{region}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &redis.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := redisOperationWaitTime(
		config.clientRedis, op, project, "Creating Instance",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("RedisInstance %q", d.Id()))
	}

	if err := d.Set("alternative_location_id", flattenRedisInstanceAlternativeLocationId(res["alternativeLocationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("authorized_network", flattenRedisInstanceAuthorizedNetwork(res["authorizedNetwork"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenRedisInstanceCreateTime(res["createTime"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("current_location_id", flattenRedisInstanceCurrentLocationId(res["currentLocationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("display_name", flattenRedisInstanceDisplayName(res["displayName"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("host", flattenRedisInstanceHost(res["host"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenRedisInstanceLabels(res["labels"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("redis_configs", flattenRedisInstanceRedisConfigs(res["redisConfigs"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("location_id", flattenRedisInstanceLocationId(res["locationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("name", flattenRedisInstanceName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memory_size_gb", flattenRedisInstanceMemorySizeGb(res["memorySizeGb"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("port", flattenRedisInstancePort(res["port"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("redis_version", flattenRedisInstanceRedisVersion(res["redisVersion"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("reserved_ip_range", flattenRedisInstanceReservedIpRange(res["reservedIpRange"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("tier", flattenRedisInstanceTier(res["tier"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceRedisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alternative_location_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, alternativeLocationIdProp)) {
		obj["alternativeLocationId"] = alternativeLocationIdProp
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	redisConfigsProp, err := expandRedisInstanceRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_configs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memory_size_gb"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, memorySizeGbProp)) {
		obj["memorySizeGb"] = memorySizeGbProp
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redis_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redisVersionProp)) {
		obj["redisVersion"] = redisVersionProp
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tier"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}
	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	if d.HasChange("memory_size_gb") {
		updateMask = append(updateMask, "memorySizeGb")
	}
	if d.HasChange("redis_configs") {
		updateMask = append(updateMask, "redisConfigs")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "PATCH", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &redis.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = redisOperationWaitTime(
		config.clientRedis, op, project, "Updating Instance",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &redis.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = redisOperationWaitTime(
		config.clientRedis, op, project, "Deleting Instance",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceRedisInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{region}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenRedisInstanceAlternativeLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceAuthorizedNetwork(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceCreateTime(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceCurrentLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceDisplayName(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceHost(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceLabels(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceRedisConfigs(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceName(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenRedisInstanceMemorySizeGb(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenRedisInstancePort(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenRedisInstanceRedisVersion(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceReservedIpRange(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceTier(v interface{}) interface{} {
	return v
}

func expandRedisInstanceAlternativeLocationId(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	fv, err := ParseNetworkFieldValue(v.(string), d, config)
	if err != nil {
		return nil, err
	}
	return fv.RelativeLink(), nil
}

func expandRedisInstanceDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceRedisConfigs(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceLocationId(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", project, region, v.(string)), nil
}

func expandRedisInstanceMemorySizeGb(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRedisVersion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReservedIpRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTier(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
