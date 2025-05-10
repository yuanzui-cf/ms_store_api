package genurl_test

import (
	"crypto/tls"
	"fmt"
	"testing"

	"ms_store_api/msstore/internal/genurl"

	"resty.dev/v3"
)

func TestGenUWPUrlWithLogs(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过需要真实网络请求的测试")
	}

	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// 尝试获取几个不同的UWP应用
	apps := []struct {
		name string
		data string
	}{
		{
			name: "Microsoft Store",
			data: `{
				"WuCategoryId": "858014f3-3934-4abe-8078-4aa193e74ca8",
				"PackageFamilyName": "Microsoft.WindowsStore_8wekyb3d8bbwe"
			}`,
		},
		{
			name: "Windows Calculator",
			data: `{
				"WuCategoryId": "c52b2d61-fa6e-4e22-b001-47b8d4b341df",
				"PackageFamilyName": "Microsoft.WindowsCalculator_8wekyb3d8bbwe"
			}`,
		},
	}

	for _, app := range apps {
		t.Run(app.name, func(t *testing.T) {
			t.Logf("\n=============== 测试应用: %s ===============", app.name)
			fileName, err := genurl.GenUWPUrl(client, app.data)

			result := fmt.Sprintf("应用: %s\n数据: %s\n结果文件名: %s\n错误: %v",
				app.name, app.data, fileName, err)
			t.Logf("\n结果汇总:\n%s", result)

			if err != nil {
				t.Errorf("获取 %s 失败: %v", app.name, err)
			} else {
				t.Logf("成功获取 %s 的文件名: %s", app.name, fileName)
			}
		})
	}
}
