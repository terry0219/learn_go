实现engine、scheduler、worker三个模块之间的通信

1. 每个worker有单独的in chan， out chan是共用的
2. workChan    chan chan engine.Request 实现chan的嵌套; 每个worker的in chan都会灌倒总的workChan
3. 将request和workChan放入队列中，不断从队列中去取任务
