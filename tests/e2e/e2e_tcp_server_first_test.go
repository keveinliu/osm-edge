package e2e

import (
	"bytes"
	"context"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	"github.com/openservicemesh/osm/pkg/constants"
	"github.com/openservicemesh/osm/tests/framework"
	. "github.com/openservicemesh/osm/tests/framework"
)

var _ = OSMDescribe("TCP server-first traffic",
	OSMDescribeInfo{
		Tier:   1,
		Bucket: 1,
	},
	func() {
		var (
			sourceNs = framework.RandomNameWithPrefix("client")
			destNs   = framework.RandomNameWithPrefix("server")
			ns       = []string{sourceNs, destNs}
		)

		It("TCP server-first traffic", func() {
			// Install OSM
			installOpts := Td.GetOSMInstallOpts()
			installOpts.EnablePermissiveMode = true
			Expect(Td.InstallOSM(installOpts)).To(Succeed())

			if sidecarClass, err := Td.GetSidecarClass(Td.OsmNamespace); err == nil {
				if strings.EqualFold(strings.ToLower(constants.SidecarClassPipy), strings.ToLower(sidecarClass)) {
					Skip("Test is only meant to be running when using envoy sidecar")
				}
			}

			// Create Test NS
			for _, n := range ns {
				Expect(Td.CreateNs(n, nil)).To(Succeed())
				Expect(Td.AddNsToMesh(true, n)).To(Succeed())
			}

			destinationPort := 80

			// Get simple pod definitions for the TCP server
			svcAccDef, podDef, svcDef, err := Td.SimplePodApp(
				SimplePodAppDef{
					PodName:     framework.RandomNameWithPrefix("server"),
					Namespace:   destNs,
					Image:       "busybox",
					Command:     []string{"nc", "-lkp", strconv.Itoa(destinationPort), "-e", "sh", "-c", "while yes; do :; done"},
					Ports:       []int{destinationPort},
					AppProtocol: constants.ProtocolTCPServerFirst,
					OS:          Td.ClusterOS,
				},
			)

			Expect(err).NotTo(HaveOccurred())

			_, err = Td.CreateServiceAccount(destNs, &svcAccDef)
			Expect(err).NotTo(HaveOccurred())
			_, err = Td.CreatePod(destNs, podDef)
			Expect(err).NotTo(HaveOccurred())
			dstSvc, err := Td.CreateService(destNs, svcDef)
			Expect(err).NotTo(HaveOccurred())

			// Expect it to be up and running in it's receiver namespace
			Expect(Td.WaitForPodsRunningReady(destNs, 120*time.Second, 1, nil)).To(Succeed())

			svcAccDef, podDef, _, err = Td.SimplePodApp(SimplePodAppDef{
				PodName:   framework.RandomNameWithPrefix("client"),
				Namespace: sourceNs,
				Command:   []string{"nc", dstSvc.Name + "." + dstSvc.Namespace, strconv.Itoa(destinationPort)},
				Image:     "busybox",
				OS:        Td.ClusterOS,
			})
			Expect(err).NotTo(HaveOccurred())
			_, err = Td.CreateServiceAccount(sourceNs, &svcAccDef)
			Expect(err).NotTo(HaveOccurred())
			_, err = Td.CreatePod(sourceNs, podDef)
			Expect(err).NotTo(HaveOccurred())

			Expect(Td.WaitForPodsRunningReady(sourceNs, 120*time.Second, 1, nil)).To(Succeed())

			Eventually(func() (string, error) {
				return getPodLogs(sourceNs, podDef.Name, podDef.Name)
			}, 5*time.Second).Should(ContainSubstring("\ny\n"), "Didn't get expected response from server")
		})
	})

// getPodLogs returns pod logs
func getPodLogs(namespace string, podName string, containerName string) (string, error) {
	podLogOptions := &corev1.PodLogOptions{
		Container: containerName,
		Follow:    false,
	}

	logStream, err := Td.Client.CoreV1().Pods(namespace).GetLogs(podName, podLogOptions).Stream(context.TODO())
	if err != nil {
		return "Error in opening stream", err
	}

	//nolint: errcheck
	//#nosec G307
	defer logStream.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(logStream)
	if err != nil {
		return "Error reading from pod logs stream", err
	}
	return buf.String(), nil
}
