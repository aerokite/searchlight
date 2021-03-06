package kube_exec

import (
	"github.com/appscode/searchlight/pkg/controller/host"
	"github.com/appscode/searchlight/test/plugin"
)

func GetTestData(objectList []*host.KubeObjectInfo) ([]plugin.TestData, error) {
	testDataList := make([]plugin.TestData, 0)
	for _, object := range objectList {
		_, objectName, namespace, err := plugin.GetKubeObjectInfo(object.Name)
		if err != nil {
			return nil, err
		}
		testData := []plugin.TestData{
			plugin.TestData{
				Data: map[string]interface{}{
					"Pod":       objectName,
					"Namespace": namespace,
					"Command":   "/bin/sh",
					"Arg":       "exit 0",
				},
				ExpectedIcingaState: 0,
			},
			plugin.TestData{
				Data: map[string]interface{}{
					"Pod":       objectName,
					"Namespace": namespace,
					"Command":   "/bin/sh",
					"Arg":       "exit 5",
				},
				ExpectedIcingaState: 2,
			},
		}
		testDataList = append(testDataList, testData...)
	}
	return testDataList, nil
}
