package fc_log

const (
	LOG_MAX_FILE_IDX = 2
	LOG_DEF_MAX_SIZE = 100 * 1024 * 1024

	LOG_INTERVAL_DAY  = 1
	LOG_INTERVAL_HOUR = 2
	LOG_INTERVAL_MIN  = 3

	LOG_OUTPUT_STD  = 0x1 << 0
	LOG_OUTPUT_FILE = 0x1 << 1
	LOG_OUTPUT_SF   = LOG_OUTPUT_STD | LOG_OUTPUT_FILE
)
