package msstore_test

import (
	"ms_store_api/msstore"
	"strings"
	"testing"
)

func TestFetchProductDetailsReal(t *testing.T) {
	t.Run("真实UWP应用测试", func(t *testing.T) {
		// 获取真实的Microsoft To-Do应用(ID: 9NBLGGH2JHXJ)
		result, err := msstore.FetchProductDetails("9NBLGGH2JHXJ")

		// 验证是否成功获取结果
		if err != nil {
			t.Fatalf("应该成功获取应用信息，但出现错误: %v", err)
		}

		if result == "" {
			t.Fatalf("返回的结果不应为空")
		}

		// 验证返回的是否为有效URL
		if !strings.Contains(result, "http") {
			t.Errorf("返回的URL应包含'http'，实际结果: %s", result)
		}

		if !strings.Contains(result, ".appx") && !strings.Contains(result, ".appxbundle") {
			t.Errorf("返回的URL应包含'.appx'或'.appxbundle'，实际结果: %s", result)
		}

		// 打印结果供参考
		t.Logf("获取到的下载链接: %s", result)
	})
}
