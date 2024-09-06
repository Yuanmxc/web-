#!/bin/bash

# 获取并杀死指定的进程
function stop_service {
    local process_name=$1
    local pids=$(tasklist //FI "IMAGENAME eq ${process_name}" //FO CSV | tail -n +2 | awk -F',' '{print $2}' | tr -d '"')

    if [[ -n "$pids" ]]; then
        for pid in $pids; do
            echo "Stopping ${process_name} service with PID $pid..."
            taskkill //PID $pid //F
        done
    else
        echo "${process_name} service not running."
    fi
}

# 停止 web 服务
stop_service "ttms_web"

# 停止 user 服务
stop_service "ttms_user"

# 停止 play 服务
stop_service "ttms_play"

# 停止 ticket 服务
stop_service "ttms_ticket"

# 停止 studio 服务
stop_service "ttms_studio"

# 停止 order 服务
stop_service "ttms_order"

echo "All services stopped."