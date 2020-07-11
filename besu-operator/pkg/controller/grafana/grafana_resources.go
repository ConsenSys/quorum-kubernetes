package grafana

import (
	// "io/ioutil"
	// "encoding/json"

	hyperledgerv1alpha1 "github.com/Sumaid/besu-kubernetes/besu-operator/pkg/apis/hyperledger/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileGrafana) grafanaConfigMapDataSources(instance *hyperledgerv1alpha1.Grafana) *corev1.ConfigMap {
	confmap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ObjectMeta.Name + "-configmap-datasources",
			Namespace: instance.Namespace,
			Labels:    r.getLabels(instance, instance.ObjectMeta.Name+"-configmap-datasources"),
		},
		Data: map[string]string{
			"prometheus.yml": `
            apiVersion: 1
            deleteDatasources:
              - name: Graphite
                orgId: 1
            datasources:
              - name: Prometheus
                type: prometheus
                access: proxy
                orgId: 1
                url: http://$PROMETHEUS_SERVICE_HOST.$NAMESPACE:9090
                password:
                user:
                database:
                basicAuth:
                basicAuthUser:
                basicAuthPassword:
                withCredentials:
                isDefault: true
                jsonData:
                  graphiteVersion: "1.1"
                  tlsAuth: false
                  tlsAuthWithCACert: false
                secureJsonData:
                  tlsCACert: "..."
                  tlsClientCert: "..."
                  tlsClientKey: "..."
                  password:
                  basicAuthPassword:
                version: 1
                editable: true`,
		},
	}
	controllerutil.SetControllerReference(instance, confmap, r.scheme)
	return confmap
}

func (r *ReconcileGrafana) grafanaConfigMapDashboardsDashboard(instance *hyperledgerv1alpha1.Grafana) *corev1.ConfigMap {
	confmap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ObjectMeta.Name + "-configmap-dashboards-dashboard",
			Namespace: instance.Namespace,
			Labels:    r.getLabels(instance, instance.ObjectMeta.Name+"-configmap-dashboards-dashboard"),
		},
		Data: map[string]string{
			"dashboard.yml": `
            apiVersion: 1
            providers:
              - name: 'Prometheus'
                orgId: 1
                folder: ''
                type: file
                disableDeletion: false
                editable: true
                options:
                  path: /etc/grafana/provisioning/dashboards`,
		},
	}
	controllerutil.SetControllerReference(instance, confmap, r.scheme)
	return confmap
}

func (r *ReconcileGrafana) grafanaConfigMapDashboardsBesu(instance *hyperledgerv1alpha1.Grafana) *corev1.ConfigMap {
	confmap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ObjectMeta.Name + "-configmap-dashboards-besu",
			Namespace: instance.Namespace,
			Labels:    r.getLabels(instance, instance.ObjectMeta.Name+"-configmap-dashboards-besu"),
		},
		Data: map[string]string{
			"besu.json": `
			{
				"__inputs": [
				{
				  "name": "Prometheus",
				  "label": "Prometheus",
				  "description": "",
				  "type": "datasource",
				  "pluginId": "prometheus",
				  "pluginName": "Prometheus"
				}
				],
				"__requires": [
				{
				  "type": "grafana",
				  "id": "grafana",
				  "name": "Grafana",
				  "version": "6.2.2"
				},
				{
				  "type": "panel",
				  "id": "graph",
				  "name": "Graph",
				  "version": ""
				},
				{
				  "type": "datasource",
				  "id": "prometheus",
				  "name": "Prometheus",
				  "version": "1.0.0"
				},
				{
				  "type": "panel",
				  "id": "table",
				  "name": "Table",
				  "version": ""
				}
				],
				"annotations": {
				  "list": [
				  {
					"builtIn": 1,
					"datasource": "-- Grafana --",
					"enable": true,
					"hide": true,
					"iconColor": "rgba(0, 211, 255, 1)",
					"name": "Annotations & Alerts",
					"type": "dashboard"
				  }
				  ]
				},
				"description": "Provides an overview of Besu nodes",
				"editable": true,
				"gnetId": 10273,
				"graphTooltip": 0,
				"id": null,
				"iteration": 1560146167177,
				"links": [],
				"panels": [
				{
				  "columns": [],
				  "fontSize": "120%",
				  "gridPos": {
					"h": 9,
					"w": 24,
					"x": 0,
					"y": 0
				  },
				  "id": 10,
				  "links": [],
				  "options": {},
				  "pageSize": null,
				  "scroll": true,
				  "showHeader": true,
				  "sort": {
					"col": 2,
					"desc": true
				  },
				  "styles": [
				  {
					"alias": "",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "job",
					"thresholds": [],
					"type": "hidden",
					"unit": "short"
				  },
				  {
					"alias": "Chain Height",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #A",
					"thresholds": [],
					"type": "number",
					"unit": "locale"
				  },
				  {
					"alias": "Total Difficulty",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #B",
					"thresholds": [],
					"type": "number",
					"unit": "sci"
				  },
				  {
					"alias": "",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Time",
					"thresholds": [],
					"type": "hidden",
					"unit": "short"
				  },
				  {
					"alias": "",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "__name__",
					"thresholds": [],
					"type": "hidden",
					"unit": "short"
				  },
				  {
					"alias": "Peer Count",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #C",
					"thresholds": [
					  ""
					],
					"type": "number",
					"unit": "locale"
				  },
				  {
					"alias": "Block Time (5m avg)",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #D",
					"thresholds": [],
					"type": "number",
					"unit": "s"
				  },
				  {
					"alias": "System",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "instance",
					"thresholds": [],
					"type": "string",
					"unit": "short",
					"valueMaps": []
				  },
				  {
					"alias": "Time Since Last Block",
					"colorMode": "value",
					"colors": [
					  "rgba(50, 172, 45, 0.97)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(245, 54, 54, 0.9)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 0,
					"mappingType": 1,
					"pattern": "Value #E",
					"thresholds": [
					  "120",
					  "240"
					],
					"type": "number",
					"unit": "dtdurations"
				  },
				  {
					"alias": "Target Chain Height",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #F",
					"thresholds": [],
					"type": "number",
					"unit": "locale"
				  },
				  {
					"alias": "Blocks Behind",
					"colorMode": "value",
					"colors": [
					  "rgba(50, 172, 45, 0.97)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(245, 54, 54, 0.9)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 2,
					"mappingType": 1,
					"pattern": "Value #G",
					"thresholds": [
					  "1",
					  "10"
					],
					"type": "number",
					"unit": "locale",
					"valueMaps": [
					{
					  "text": "Yes",
					  "value": "1"
					},
					{
					  "text": "No",
					  "value": "0"
					}
					]
				  },
				  {
					"alias": "% Peer Limit Used",
					"colorMode": "value",
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"dateFormat": "YYYY-MM-DD HH:mm:ss",
					"decimals": 0,
					"mappingType": 1,
					"pattern": "Value #H",
					"thresholds": [
					  "0.25",
					  "0.75"
					],
					"type": "number",
					"unit": "percentunit"
				  },
				  {
					"alias": "",
					"colorMode": null,
					"colors": [
					  "rgba(245, 54, 54, 0.9)",
					  "rgba(237, 129, 40, 0.89)",
					  "rgba(50, 172, 45, 0.97)"
					],
					"decimals": 2,
					"pattern": "/.*/",
					"thresholds": [],
					"type": "number",
					"unit": "short"
				  }
				  ],
				  "targets": [
				  {
					"expr": "sum by (instance) (ethereum_blockchain_height{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"interval": "",
					"intervalFactor": 1,
					"legendFormat": "Chain Height",
					"refId": "A"
				  },
				  {
					"expr": "sum by (instance) (ethereum_best_known_block_number{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"refId": "F"
				  },
				  {
					"expr": "sum by (instance) (ethereum_best_known_block_number{instance=~\"$system\"} - ethereum_blockchain_height{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"refId": "G"
				  },
				  {
					"expr": "sum by (instance) (besu_blockchain_difficulty_total{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"legendFormat": "Total Difficulty",
					"refId": "B"
				  },
				  {
					"expr": "sum by (instance) (1/rate(ethereum_blockchain_height{instance=~\"$system\"}[5m]))",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"refId": "D"
				  },
				  {
					"expr": "sum by (instance) (time() - besu_blockchain_chain_head_timestamp{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"refId": "E"
				  },
				  {
					"expr": "sum by (instance) (ethereum_peer_count{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"legendFormat": "Peer Count",
					"refId": "C"
				  },
				  {
					"expr": "sum by (instance) (ethereum_peer_count{instance=~\"$system\"} / ethereum_peer_limit{instance=~\"$system\"})",
					"format": "table",
					"instant": true,
					"intervalFactor": 1,
					"refId": "H"
				  }
				  ],
				  "title": "Overview",
				  "transform": "table",
				  "type": "table"
				},
				{
				  "aliasColors": {},
				  "bars": false,
				  "dashLength": 10,
				  "dashes": false,
				  "fill": 1,
				  "gridPos": {
					"h": 14,
					"w": 12,
					"x": 0,
					"y": 9
				  },
				  "id": 12,
				  "legend": {
					"alignAsTable": true,
					"avg": true,
					"current": true,
					"max": true,
					"min": true,
					"rightSide": false,
					"show": true,
					"total": false,
					"values": true
				  },
				  "lines": true,
				  "linewidth": 1,
				  "links": [],
				  "nullPointMode": "null",
				  "options": {},
				  "percentage": false,
				  "pointradius": 5,
				  "points": false,
				  "renderer": "flot",
				  "seriesOverrides": [],
				  "spaceLength": 10,
				  "stack": false,
				  "steppedLine": false,
				  "targets": [
				  {
					"expr": "1/rate(ethereum_blockchain_height{instance=~\"$system\"}[5m])",
					"format": "time_series",
					"interval": "",
					"intervalFactor": 1,
					"legendFormat": "{{instance}}",
					"refId": "A"
				  },
				  {
					"expr": "",
					"format": "time_series",
					"intervalFactor": 1,
					"refId": "B"
				  }
				  ],
				  "thresholds": [],
				  "timeFrom": null,
				  "timeRegions": [],
				  "timeShift": null,
				  "title": "Block Time",
				  "tooltip": {
					"shared": true,
					"sort": 0,
					"value_type": "individual"
				  },
				  "type": "graph",
				  "xaxis": {
					"buckets": null,
					"mode": "time",
					"name": null,
					"show": true,
					"values": []
				  },
				  "yaxes": [
				  {
					"format": "s",
					"label": null,
					"logBase": 10,
					"max": null,
					"min": null,
					"show": true
				  },
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  }
				  ],
				  "yaxis": {
					"align": false,
					"alignLevel": null
				  }
				},
				{
				  "aliasColors": {},
				  "bars": false,
				  "dashLength": 10,
				  "dashes": false,
				  "fill": 1,
				  "gridPos": {
					"h": 14,
					"w": 12,
					"x": 12,
					"y": 9
				  },
				  "id": 13,
				  "legend": {
					"alignAsTable": true,
					"avg": true,
					"current": true,
					"max": true,
					"min": true,
					"rightSide": false,
					"show": true,
					"total": false,
					"values": true
				  },
				  "lines": true,
				  "linewidth": 1,
				  "links": [],
				  "nullPointMode": "null",
				  "options": {},
				  "percentage": false,
				  "pointradius": 5,
				  "points": false,
				  "renderer": "flot",
				  "seriesOverrides": [],
				  "spaceLength": 10,
				  "stack": false,
				  "steppedLine": false,
				  "targets": [
				  {
					"expr": "ethereum_best_known_block_number{instance=~\"$system\"} - ethereum_blockchain_height{instance=~\"$system\"}",
					"format": "time_series",
					"interval": "",
					"intervalFactor": 1,
					"legendFormat": "{{instance}}",
					"refId": "A"
				  }
				  ],
				  "thresholds": [],
				  "timeFrom": null,
				  "timeRegions": [],
				  "timeShift": null,
				  "title": "Blocks Behind",
				  "tooltip": {
					"shared": true,
					"sort": 0,
					"value_type": "individual"
				  },
				  "type": "graph",
				  "xaxis": {
					"buckets": null,
					"mode": "time",
					"name": null,
					"show": true,
					"values": []
				  },
				  "yaxes": [
				  {
					"format": "locale",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  },
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  }
				  ],
				  "yaxis": {
					"align": false,
					"alignLevel": null
				  }
				},
				{
				  "aliasColors": {},
				  "bars": false,
				  "dashLength": 10,
				  "dashes": false,
				  "fill": 1,
				  "gridPos": {
					"h": 12,
					"w": 24,
					"x": 0,
					"y": 23
				  },
				  "id": 6,
				  "legend": {
					"avg": false,
					"current": false,
					"max": false,
					"min": false,
					"show": true,
					"total": false,
					"values": false
				  },
				  "lines": true,
				  "linewidth": 1,
				  "links": [],
				  "nullPointMode": "null",
				  "options": {},
				  "percentage": false,
				  "pointradius": 5,
				  "points": false,
				  "renderer": "flot",
				  "seriesOverrides": [],
				  "spaceLength": 10,
				  "stack": false,
				  "steppedLine": false,
				  "targets": [
				  {
					"expr": "irate(process_cpu_seconds_total{instance=~\"$system\"}[1m])",
					"format": "time_series",
					"intervalFactor": 1,
					"legendFormat": "CPU Time IRate [{{instance}}]",
					"refId": "A"
				  }
				  ],
				  "thresholds": [],
				  "timeFrom": null,
				  "timeRegions": [],
				  "timeShift": null,
				  "title": "CPU",
				  "tooltip": {
					"shared": true,
					"sort": 0,
					"value_type": "individual"
				  },
				  "type": "graph",
				  "xaxis": {
					"buckets": null,
					"mode": "time",
					"name": null,
					"show": true,
					"values": []
				  },
				  "yaxes": [
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  },
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  }
				  ],
				  "yaxis": {
					"align": false,
					"alignLevel": null
				  }
				},
				{
				  "aliasColors": {},
				  "bars": false,
				  "dashLength": 10,
				  "dashes": false,
				  "datasource": "Prometheus",
				  "fill": 1,
				  "gridPos": {
					"h": 12,
					"w": 12,
					"x": 0,
					"y": 35
				  },
				  "id": 8,
				  "legend": {
					"avg": false,
					"current": false,
					"max": false,
					"min": false,
					"show": true,
					"total": false,
					"values": false
				  },
				  "lines": true,
				  "linewidth": 1,
				  "links": [],
				  "nullPointMode": "null",
				  "options": {},
				  "percentage": false,
				  "pointradius": 5,
				  "points": false,
				  "renderer": "flot",
				  "seriesOverrides": [],
				  "spaceLength": 10,
				  "stack": false,
				  "steppedLine": false,
				  "targets": [
				  {
					"expr": "rate(jvm_gc_collection_seconds_sum{instance=~\"$system\"}[1m])",
					"format": "time_series",
					"interval": "",
					"intervalFactor": 5,
					"legendFormat": "{{gc}} [{{instance}}]",
					"metric": "jvm_gc_collection_seconds_sum",
					"refId": "A",
					"step": 10
				  }
				  ],
				  "thresholds": [],
				  "timeFrom": null,
				  "timeRegions": [],
				  "timeShift": null,
				  "title": "GC time",
				  "tooltip": {
					"shared": true,
					"sort": 0,
					"value_type": "individual"
				  },
				  "type": "graph",
				  "xaxis": {
					"buckets": null,
					"mode": "time",
					"name": null,
					"show": true,
					"values": []
				  },
				  "yaxes": [
				  {
					"decimals": null,
					"format": "percentunit",
					"label": null,
					"logBase": 1,
					"max": "1",
					"min": "0",
					"show": true
				  },
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  }
				  ],
				  "yaxis": {
					"align": false,
					"alignLevel": null
				  }
				},
				{
				  "aliasColors": {},
				  "bars": false,
				  "dashLength": 10,
				  "dashes": false,
				  "fill": 1,
				  "gridPos": {
					"h": 12,
					"w": 12,
					"x": 12,
					"y": 35
				  },
				  "id": 4,
				  "legend": {
					"alignAsTable": true,
					"avg": true,
					"current": true,
					"max": true,
					"min": true,
					"show": true,
					"total": false,
					"values": true
				  },
				  "lines": true,
				  "linewidth": 1,
				  "links": [],
				  "nullPointMode": "null",
				  "options": {},
				  "percentage": false,
				  "pointradius": 5,
				  "points": false,
				  "renderer": "flot",
				  "seriesOverrides": [],
				  "spaceLength": 10,
				  "stack": false,
				  "steppedLine": false,
				  "targets": [
				  {
					"expr": "jvm_memory_bytes_used{instance=~\"$system\", area=\"heap\"} + ignoring(area) jvm_memory_bytes_used{instance=~\"$system\", area=\"nonheap\"}",
					"format": "time_series",
					"intervalFactor": 5,
					"legendFormat": "{{instance}}",
					"refId": "A"
				  }
				  ],
				  "thresholds": [],
				  "timeFrom": null,
				  "timeRegions": [],
				  "timeShift": null,
				  "title": "Memory Used",
				  "tooltip": {
					"shared": true,
					"sort": 0,
					"value_type": "individual"
				  },
				  "type": "graph",
				  "xaxis": {
					"buckets": null,
					"mode": "time",
					"name": null,
					"show": true,
					"values": []
				  },
				  "yaxes": [
				  {
					"format": "decbytes",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  },
				  {
					"format": "short",
					"label": null,
					"logBase": 1,
					"max": null,
					"min": null,
					"show": true
				  }
				  ],
				  "yaxis": {
					"align": false,
					"alignLevel": null
				  }
				}
				],
				"refresh": "10s",
				"schemaVersion": 18,
				"style": "dark",
				"tags": [
				  "besu",
				  "ethereum"
				],
				"templating": {
				  "list": [
				  {
					"allValue": null,
					"current": {},
					"datasource": "Prometheus",
					"definition": "ethereum_blockchain_height",
					"hide": 0,
					"includeAll": true,
					"label": "System",
					"multi": true,
					"name": "system",
					"options": [],
					"query": "ethereum_blockchain_height",
					"refresh": 2,
					"regex": "/instance=\"([^\"]*)\"/",
					"skipUrlSync": false,
					"sort": 5,
					"tagValuesQuery": "",
					"tags": [],
					"tagsQuery": "",
					"type": "query",
					"useTags": false
				  }
				  ]
				},
				"time": {
				  "from": "now-12h",
				  "to": "now"
				},
				"timepicker": {
				  "refresh_intervals": [
					"5s",
					"10s",
					"30s",
					"1m",
					"5m",
					"15m",
					"30m",
					"1h",
					"2h",
					"1d"
				  ],
				  "time_options": [
					"5m",
					"15m",
					"1h",
					"6h",
					"12h",
					"24h",
					"2d",
					"7d",
					"30d"
				  ]
				},
				"timezone": "",
				"title": "Besu Overview",
				"uid": "XE4V0WGZz",
				"version": 1
			  }`,
		},
	}
	controllerutil.SetControllerReference(instance, confmap, r.scheme)
	return confmap
}

func (r *ReconcileGrafana) grafanaService(instance *hyperledgerv1alpha1.Grafana) *corev1.Service {

	serv := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ObjectMeta.Name,
			Namespace: instance.Namespace,
			Labels:    r.getLabels(instance, instance.ObjectMeta.Name),
		},
		Spec: corev1.ServiceSpec{
			Type:     "NodePort",
			Selector: r.getLabels(instance, instance.ObjectMeta.Name),
			Ports: []corev1.ServicePort{
				{
					Name:       instance.ObjectMeta.Name,
					Protocol:   "TCP",
					Port:       int32(3000),
					TargetPort: intstr.FromInt(int(3000)),
					NodePort:   int32(30030),
				},
			},
		},
	}
	controllerutil.SetControllerReference(instance, serv, r.scheme)
	return serv
}

func (r *ReconcileGrafana) grafanaDeployment(instance *hyperledgerv1alpha1.Grafana) *appsv1.Deployment {

	var replicas int32
	replicas = 1
	depl := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.ObjectMeta.Name,
			Namespace: instance.Namespace,
			Labels:    r.getLabels(instance, instance.ObjectMeta.Name),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: r.getLabels(instance, instance.ObjectMeta.Name),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: r.getLabels(instance, instance.ObjectMeta.Name),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						corev1.Container{
							Name:            instance.ObjectMeta.Name,
							Image:           "grafana/grafana:6.2.5",
							ImagePullPolicy: "IfNotPresent",
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("100m"),
									corev1.ResourceMemory: resource.MustParse("256Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resource.MustParse("500m"),
									corev1.ResourceMemory: resource.MustParse("512Mi"),
								},
								// Requests: corev1.ResourceList{
								// 	corev1.ResourceCPU:    resource.MustParse(instance.Spec.Resources.CPURequest),
								// 	corev1.ResourceMemory: resource.MustParse(instance.Spec.Resources.MemRequest),
								// },
								// Limits: corev1.ResourceList{
								// 	corev1.ResourceCPU:    resource.MustParse(instance.Spec.Resources.CPULimit),
								// 	corev1.ResourceMemory: resource.MustParse(instance.Spec.Resources.MemLimit),
								// },
							},
							Env: []corev1.EnvVar{
								{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
								{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								{
									Name:  "GF_SECURITY_ADMIN_USER",
									Value: "admin",
								},
								{
									Name:  "GF_SECURITY_ADMIN_PASSWORD",
									Value: "password",
								},
								{
									Name:  "PROMETHEUS_SERVICE_HOST",
									Value: "prometheus",
								},
								{
									Name:  "NAMESPACE",
									Value: instance.Namespace,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								corev1.VolumeMount{
									Name:      instance.ObjectMeta.Name + "-configmap-datasources",
									MountPath: "/etc/grafana/provisioning/datasources/prometheus.yml",
									SubPath:   "prometheus.yml",
									ReadOnly:  true,
								},
								corev1.VolumeMount{
									Name:      instance.ObjectMeta.Name + "-configmap-dashboards-dashboard",
									MountPath: "/etc/grafana/provisioning/dashboards/dashboard.yml",
									SubPath:   "dashboard.yml",
									ReadOnly:  true,
								},
								corev1.VolumeMount{
									Name:      instance.ObjectMeta.Name + "-configmap-dashboards-besu",
									MountPath: "/etc/grafana/provisioning/dashboards/besu.json",
									SubPath:   "besu.json",
									ReadOnly:  true,
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: int32(3000),
									Protocol:      "TCP",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						corev1.Volume{
							Name: instance.ObjectMeta.Name + "-configmap-datasources",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: instance.ObjectMeta.Name + "-configmap-datasources",
									},
								},
							},
						},
						corev1.Volume{
							Name: instance.ObjectMeta.Name + "-configmap-dashboards-dashboard",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: instance.ObjectMeta.Name + "-configmap-dashboards-dashboard",
									},
								},
							},
						},
						corev1.Volume{
							Name: instance.ObjectMeta.Name + "-configmap-dashboards-besu",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: instance.ObjectMeta.Name + "-configmap-dashboards-besu",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	controllerutil.SetControllerReference(instance, depl, r.scheme)
	return depl
}

func (r *ReconcileGrafana) getLabels(instance *hyperledgerv1alpha1.Grafana, name string) map[string]string {
	labels := make(map[string]string)
	labels["app"] = name
	return labels
}
