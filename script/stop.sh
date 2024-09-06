#!/bin/bash

# 停止 web 服务
web_pid=$(pgrep -f "ttms_web")
if [[ -n "$web_pid" ]]; then
    echo "Stopping web service..."
    kill $web_pid
fi

# 停止 user 服务
user_pid=$(pgrep -f "ttms_user")
if [[ -n "$user_pid" ]]; then
    echo "Stopping user service..."
    kill $user_pid
fi

# 停止 play 服务
play_pid=$(pgrep -f "ttms_play")
if [[ -n "$play_pid" ]]; then
    echo "Stopping play service..."
    kill $play_pid
fi

# 停止 ticket 服务
ticket_pid=$(pgrep -f "ttms_ticket")
if [[ -n "$ticket_pid" ]]; then
    echo "Stopping ticket service..."
    kill $ticket_pid
fi

# 停止 studio 服务
studio_pid=$(pgrep -f "ttms_studio")
if [[ -n "$studio_pid" ]]; then
    echo "Stopping studio service..."
    kill $studio_pid
fi

# 停止 order 服务
order_pid=$(pgrep -f "ttms_order")
if [[ -n "$order_pid" ]]; then
    echo "Stopping order service..."
    kill $order_pid
fi

echo "All services stopped."
