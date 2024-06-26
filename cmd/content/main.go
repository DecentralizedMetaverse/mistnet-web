package main

import (
	"fmt"
	"net/http"
	"os"

	"mistlib/internal/content"
	"mistlib/internal/content/ipfs"
)

func main() {

	if len(os.Args) > 1 {
		// コマンドライン引数が存在する場合、コマンドラインモードで実行
		content.RunCommand(os.Args[1], os.Args[2:])
	} else {
		ipfs.StartDaemon()
		// 引数がない場合、HTTPサーバーを起動
		http.HandleFunc("/init", content.HandleInitAPI)
		http.HandleFunc("/switch", content.HandleSwitchAPI)
		http.HandleFunc("/get", content.HandleGetAPI)
		http.HandleFunc("/add", content.HandleAddAPI)
		http.HandleFunc("/remove", content.HandleRemoveAPI)
		http.HandleFunc("/set-password", content.HandleSetPasswordAPI)
		http.HandleFunc("/cat", content.HandleCatAPI)
		http.HandleFunc("/get-world-cid", content.HandleGetWorldCIDAPI)
		http.HandleFunc("/download-world", content.HandleDownloadWorldAPI)
		http.HandleFunc("/get-world-info", content.HandleGetWorldInfoAPI)
		http.HandleFunc("/unpack", content.HandleUnpackAPI)
		http.HandleFunc("/set-custom-data", content.HandleSetCustomDataAPI)
		http.HandleFunc("/set-parent", content.HandleSetParentAPI)
		http.HandleFunc("/update", content.HandleUpdateAPI)

		fmt.Println("Starting server on :8081")
		http.ListenAndServe(":8081", nil)

	}
}
