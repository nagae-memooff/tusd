#!/bin/sh
#
# Init Script to run tusd in daemon mode at boot time.
# chkconfig: - 67 67
# description: tusd daemon
# processname: tusd
#
# Run "/sbin/chkconfig --add mx_tusd" to add the Run levels.
# This will setup the symlinks and set the process to run at boot.

# Source function library.
. /etc/rc.d/init.d/functions

if [ -f /opt/ewhine/etc/init.d/mx_env ]; then
  . /opt/ewhine/etc/init.d/mx_env
fi

USER=ewhine
DIR="/opt/ewhine/bin/"
DAEMON="$DIR/tusd"
PIDFILE=/home/ewhine/var/run/tusd.pid
LOGFILE=/home/ewhine/var/log/tusd.log
ulimit -n 100000

STORE_DIR="/opt/ewhine/deploy/ewhine_NB/efiles/contents/tusd"
ARGS="-dir $STORE_DIR"

RETVAL=0
prog="tusd"

if [ ! -f $PIDFILE ]; then
  touch $PIDFILE
  chown ${USER}:${USER} $PIDFILE
fi

start () {
  PID=`pidof tusd`
  if [ $? -eq 0 ]; then
    echo "$prog has already been started!"
    exit 0;
  fi
	mkdir -p $STORE_DIR

  echo -n "Starting $prog: "
  #start-stop-daemon --start --background --chuid $USER --chdir $DIR --make-pidfile --pidfile $PIDFILE --exec $DAEMON
  start-stop-daemon --start --background --chuid $USER --chdir $DIR --make-pidfile --pidfile $PIDFILE --startas /bin/bash -- -c "exec $DAEMON $ARGS >> $LOGFILE 2>&1"

  

  RETVAL=$?
  if [ $RETVAL -eq 0 ]; then
    sleep 1
    chown ${USER}:${USER} $PIDFILE
    success
  else
    failure
  fi
  echo
  return $RETVAL
}

stop () {
  status2()
  if [ $? -ne 0 ]; then
    echo "$prog is not running!"
    exit 0;
  fi

  echo -n "Stopping $prog: "
  killall $DAEMON > /dev/null 2>&1
  RETVAL=$?
  if [ $RETVAL -eq 0 ] ; then
    success
  else
    failure
  fi
  echo
  return $RETVAL
}

status2 () {
  pid=`pidof $DAEMON`
  if [ $? -eq 0 ]; then
    echo "$prog is running with pid $pid"
    return 0;
  else
    echo "$prog is not running!"
    return 1;
  fi
  return 2
}

restart () {
  stop
  start
}

# See how we were called.
case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  status)
    status2
    ;;
  restart|reload)
    restart
    ;;
  *)
    echo "Usage: $0 {start|stop|status|restart|reload}"
    exit 1
esac

exit $?
