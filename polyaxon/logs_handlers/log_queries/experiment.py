from django.conf import settings

from polyaxon_k8s.manager import K8SManager

from constants.k8s_jobs import EXPERIMENT_JOB_NAME_FORMAT
from logs_handlers.log_queries import base
from logs_handlers.utils import safe_log_experiment
from schemas.tasks import TaskType


def stream_logs(experiment):
    pod_id = EXPERIMENT_JOB_NAME_FORMAT.format(
        task_type=TaskType.MASTER,  # We default to master
        task_idx=0,
        experiment_uuid=experiment.uuid.hex)
    k8s_manager = K8SManager(namespace=settings.K8S_NAMESPACE, in_cluster=True)
    return base.stream_logs(k8s_manager=k8s_manager,
                            pod_id=pod_id,
                            container_job_name=settings.CONTAINER_NAME_EXPERIMENT_JOB)


def process_logs(experiment, temp=True):
    pod_id = EXPERIMENT_JOB_NAME_FORMAT.format(
        task_type=TaskType.MASTER,  # We default to master
        task_idx=0,
        experiment_uuid=experiment.uuid.hex)
    k8s_manager = K8SManager(namespace=settings.K8S_NAMESPACE, in_cluster=True)
    log_lines = base.process_logs(k8s_manager=k8s_manager,
                                  pod_id=pod_id,
                                  container_job_name=settings.CONTAINER_NAME_EXPERIMENT_JOB)

    safe_log_experiment(experiment_name=experiment.unique_name, log_lines=log_lines, temp=temp)
