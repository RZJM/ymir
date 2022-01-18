from typing import Dict, Optional

from pydantic import BaseModel, Field


class TaskStateExtra(BaseModel):
    user_id: str


class TaskStatePercent(BaseModel):
    task_id: str
    timestamp: int = Field(gt=0)
    percent: float = Field(ge=0, le=1)
    state: int
    state_code: int
    state_message: Optional[str]
    stack_error_info: Optional[str]


class TaskState(BaseModel):
    task_extra_info: TaskStateExtra
    percent_result: TaskStatePercent


TaskStateDict = Dict[str, TaskState]


# data models: resp
class EventResp(BaseModel):
    return_code: int
    return_msg: str