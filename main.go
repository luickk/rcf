package main

import (
 rcf_cc_node "robot-communication-framework/rcf_cc_node"
)

func main() {
  rcf_cc_node.Init(200)

  rcf_cc_node.Create_cctopic("test")
}
