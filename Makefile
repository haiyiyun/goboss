PWD=`pwd`
BIN=$(PWD)/bin
GO_DEBUG=",debug"
GO_PORT_GOBOSS="8080"

#==========================================
install_goboss:
	GOBIN=$(BIN) go install ./cmd/goboss/...

run_goboss:
ifeq (bin/pid.goboss, $(wildcard bin/pid.goboss))
	@echo 'Process exists PID: ' && cat bin/pid.goboss
else
	@mkdir -p $(BIN)/run/logs
	cd $(BIN) && `GO_PORT=$(GO_PORT_GOBOSS) GO_DEBUG=$(GO_DEBUG) nohup ./goboss > ./run/logs/nohup_goboss.log 2>&1 & echo $$! >./pid.goboss`
endif

status_goboss:
	@cd $(BIN) && pgrep -aF ./pid.goboss

kill_goboss:
	@cd $(BIN) && pkill -9 -eF ./pid.goboss && rm ./pid.goboss

log_goboss:
	@tail -f $(BIN)/run/logs/goboss.log -f $(BIN)/run/logs/nohup_goboss.log
#==========================================

install:
	make install_goboss
#==========================================

runall:
	make run_goboss
#==========================================

statusall:
	make status_goboss
#==========================================

killall:
	make kill_goboss
#==========================================

logall:
	make log_goboss
#==========================================

status:
	make statusall | grep -e ^[0-9]
#==========================================

rerunall:
	make killall
	make runall
#==========================================

reinstall:
	make killall
	make install
	make runall
#==========================================
clean:
	rm -rf $(BIN)/