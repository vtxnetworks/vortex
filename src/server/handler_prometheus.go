package server

import (
	"github.com/linkernetworks/vortex/src/entity"
	response "github.com/linkernetworks/vortex/src/net/http"
	"github.com/linkernetworks/vortex/src/net/http/query"
	pc "github.com/linkernetworks/vortex/src/prometheuscontroller"
	"github.com/linkernetworks/vortex/src/web"
)

func listContainerMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	queryLabels := map[string]string{}

	if node, ok := query.Str("node"); ok {
		queryLabels["node"] = node
	}

	if namespace, ok := query.Str("namespace"); ok {
		queryLabels["namespace"] = namespace
	}

	if pod, ok := query.Str("pod"); ok {
		queryLabels["pod"] = pod
	}

	containerNameList, err := pc.ListContainerName(sp, queryLabels)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	containerList := map[string]entity.ContainerMetrics{}
	for _, containerName := range containerNameList {
		container, err := pc.GetContainer(sp, containerName)
		if err != nil {
			response.InternalServerError(req.Request, resp.ResponseWriter, err)
			return
		}
		containerList[containerName] = container
	}

	resp.WriteEntity(containerList)
}

func getContainerMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	container, err := pc.GetContainer(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(container)
}

func listPodMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	queryLabels := map[string]string{}

	if node, ok := query.Str("node"); ok {
		queryLabels["node"] = node
	}

	if namespace, ok := query.Str("namespace"); ok {
		queryLabels["namespace"] = namespace
	}

	if controller, ok := query.Str("controller"); ok {
		queryLabels["created_by_kind"] = "ReplicaSet"
		queryLabels["created_by_name"] = controller + ".*"
	}

	podNameList, err := pc.ListPodName(sp, queryLabels)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	podList := map[string]entity.PodMetrics{}
	for _, podName := range podNameList {
		pod, err := pc.GetPod(sp, podName)
		if err != nil {
			response.InternalServerError(req.Request, resp.ResponseWriter, err)
			return
		}
		podList[podName] = pod
	}

	resp.WriteEntity(podList)
}

func getPodMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	pod, err := pc.GetPod(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(pod)
}

func listServiceMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	queryLabels := map[string]string{}

	if namespace, ok := query.Str("namespace"); ok {
		queryLabels["namespace"] = namespace
	}

	serviceNameList, err := pc.ListServiceName(sp, queryLabels)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	serviceList := map[string]entity.ServiceMetrics{}
	for _, serviceName := range serviceNameList {
		service, err := pc.GetService(sp, serviceName)
		if err != nil {
			response.InternalServerError(req.Request, resp.ResponseWriter, err)
			return
		}
		serviceList[serviceName] = service
	}

	resp.WriteEntity(serviceList)
}

func getServiceMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	service, err := pc.GetService(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(service)
}

func listControllerMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	queryLabels := map[string]string{}

	if namespace, ok := query.Str("namespace"); ok {
		queryLabels["namespace"] = namespace
	}

	controllerNameList, err := pc.ListControllerName(sp, queryLabels)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	controllerList := map[string]entity.ControllerMetrics{}
	for _, controllerName := range controllerNameList {
		controller, err := pc.GetController(sp, controllerName)
		if err != nil {
			response.InternalServerError(req.Request, resp.ResponseWriter, err)
			return
		}
		controllerList[controllerName] = controller
	}

	resp.WriteEntity(controllerList)
}

func getControllerMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	controller, err := pc.GetController(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(controller)
}

func listNodeMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response

	query := query.New(req.Request.URL.Query())
	queryLabels := map[string]string{}

	nodeNameList, err := pc.ListNodeName(sp, queryLabels)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	nodeList := map[string]entity.NodeMetrics{}
	for _, nodeName := range nodeNameList {
		node, err := pc.GetNode(sp, nodeName)
		if err != nil {
			response.InternalServerError(req.Request, resp.ResponseWriter, err)
			return
		}
		nodeList[nodeName] = node
	}

	resp.WriteEntity(nodeList)
}

func listNodeNicsMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	nicList, err := pc.ListNodeNICs(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(nicList)
}

func getNodeMetricsHandler(ctx *web.Context) {
	sp, req, resp := ctx.ServiceProvider, ctx.Request, ctx.Response
	id := req.PathParameter("id")

	node, err := pc.GetNode(sp, id)
	if err != nil {
		response.InternalServerError(req.Request, resp.ResponseWriter, err)
		return
	}

	resp.WriteEntity(node)
}
