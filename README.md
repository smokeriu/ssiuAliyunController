为了防止自用服务器被挖矿。由于家里网络没有稳定的公网ip，在ip发生变动且无法访问后，需要更新阿里云某一安全组的允许ip。
保持同样的端口、描述等，只更新组内的允许的源IP地址。

# TODO List
- Handle error
- Use goroutine
- More flexible use
- Run on Docker via Synology