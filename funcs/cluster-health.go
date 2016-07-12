package funcs

var clusterHealth = map[string]string{
	"status":                    "GAUGE",
	"timed_out":                 "GAUGE",
	"number_of_nodes":           "GAUGE",
	"number_of_data_nodes":      "GAUGE",
	"active_primary_shards":     "GAUGE",
	"active_shards":             "GAUGE",
	"relocating_shards":         "GAUGE",
	"initializing_shards":       "GAUGE",
	"unassigned_shards":         "GAUGE",
	"delayed_unassigned_shards": "GAUGE",
	"number_of_pending_tasks":   "GAUGE",
	"number_of_in_flight_fetch": "GAUGE",
}
