#!/usr/bin/env bash

cluster_addr=$1
cluster_name=$2
sst_request_opts=$3
log_file=$4
socat_opts=$5

# Here, my_ip is the IP of the host on which this script is running
# And opts has 3 parts.
#   part-01: The name of the method/script the Galera Arbitrator (garbd) uses in a State Snapshot Transfer (SST)
#            For xtrabackup, the method is "xtrabackup-v2"
#   part-02: The port at which Galera Arbitrator (garbd) listen to take SST
#            For xtrabackup, the port is 4444
#   part-03: Suffix of sst_request string
#            For xtrabackup, the suffix is "/xtrabackup_sst//1"
#
# So, finally the sst_request_string is of following format:
#     "<method>:<ip>:<port><suffix>"
function get_sst_request_string() {
    my_ip=$1
    opts=($(echo "$2" | sed -e "s/,/ /g"))

    sst_request_method=${opts[0]}
    sst_request_port=${opts[1]}
    sst_request_suffix=${opts[2]}

    printf "%s:%s:%s%s" $sst_request_method $my_ip $sst_request_port $sst_request_suffix
}

# Command $(hostname -I) returns a space separated IP list. We need only the first one.
myips=$(hostname -I)
first=${myips%% *}
sst_request_string=$(get_sst_request_string $first $sst_request_opts)

# Start backup procedure
echo "" > $log_file
timeout -k 25 20 \
  garbd \
    --address="$cluster_addr" \
    --group="$cluster_name" \
    --sst="$sst_request_string" \
    --log="$log_file"

SOCAT_OPTS=$socat_opts

# the first socat run does not give the streams. So, we need to run socat again
socat -u "$SOCAT_OPTS" stdio > /dev/null
socat -u "$SOCAT_OPTS" stdio
