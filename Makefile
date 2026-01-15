DATA_DIR := data
ID_FILE := $(DATA_DIR)/user_ids.txt

.PHONY: all gen bloom clean

all: gen bloom

gen:
	@mkdir -p $(DATA_DIR)
	go run gen_ids.go -count 1000000 -out $(ID_FILE)

bloom:
	go run bloom_filter.go -in $(ID_FILE) -false-pos 0.001 -queries 100000

clean:
	rm -f $(ID_FILE)
