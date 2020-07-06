# rabbitmq-metrics-emitter-release

This repo contains the code and boshrelease for the rabbitmq-metrics-emitter.

When deployed together with the [https://github.com/pivotal-cf/cf-rabbitmq-multitenant-broker-release](multitenant rabbitmq service-broker) it will emit the metrics needed to configure an app autoscaling policy to react to queue depth (as measured by the `messages_ready` metric).

Metrics are emmitted under the name `<queue-name>-<metric-name>` currently the `messages_ready` is the only one being recorded

Here is an example autoscaler policy that would react to the number of unprocessed messages in the dummy-queue. Note that the policy can only reference queues created on the service instance bound to the target app.

```
{
	"instance_min_count": 1,
	"instance_max_count": 3,
	"scaling_rules": [
    {
      "metric_type": "dummy_queue_messages_ready",
      "threshold": 1,
      "operator": ">",
      "adjustment": "+1"
    },
    {
      "metric_type": "dummy_queue_messages_ready",
      "threshold": 2,
      "operator": "<",
      "adjustment": "-1"
    }
  ]
}
```
