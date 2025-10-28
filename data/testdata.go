package data

const TestLogData = `
{
  "@timestamp": "2025-10-23T15:43:33.01964036+07:00",
  "agent": {
    "id": "099",
    "ip": "10.80.140.25",
    "name": "hybrid-sensor"
  },
  "cluster": {
    "name": "wazuh-server-cluster",
    "node": "wazuh-server-3"
  },
  "decoder": {
    "name": "multi-source"
  },
  "os": {
    "type": "linux",
    "version": "5.15.0-116-generic"
  },
  "rule": {
    "description": "Mixed event test rule",
    "id": "90001",
    "level": 3,
    "tags": ["test", "hybrid", "validation"]
  },
  "linux_log": {
    "full_log": "Oct 22 15:17:16 DB systemd: vault.service: main process exited, code=exited, status=1/FAILURE",
    "predecoder": {
      "hostname": "DB",
      "program_name": "systemd",
      "timestamp": "Oct 22 15:17:16"
    },
    "metadata": {
      "facility": "daemon",
      "priority": "error",
      "pid": 3124,
      "custom_flag": true
    }
  },
  "windows_event": {
    "system": {
      "providerName": "Microsoft-Windows-Sysmon",
      "eventID": "6",
      "channel": "Microsoft-Windows-Sysmon/Operational",
      "computer": "AD-DC.signad-dc.local",
      "systemTime": "2025-10-22T08:17:15.721633100Z"
    },
    "eventdata": {
      "imageLoaded": "C:\\\\Windows\\\\System32\\\\drivers\\\\fortiwf2.sys",
      "hashes": {
        "SHA1": "4C3AC7D7E3771450BEA763F51EEBD5843F77F09C",
        "MD5": "6EE1DB36B89202226C98F1D8B8A1CA01",
        "SHA256": "999765DA696087E060A8FB7F262DB659A4909768EDEDDBD3595F136CF59CCA46"
      },
      "signed": true,
      "signature": "Fortinet, Inc.",
      "signatureStatus": "Valid"
    }
  },
  "cloud_metadata": {
    "provider": "aws",
    "region": "ap-southeast-1",
    "instance_id": "i-0abc1234def56789",
    "tags": {
      "Environment": "staging",
      "Service": "collector"
    }
  },
  "network_activity": [
    {
      "protocol": "TCP",
      "src_ip": "10.80.130.5",
      "dst_ip": "10.80.150.10",
      "src_port": 443,
      "dst_port": 52345,
      "action": "allow"
    },
    {
      "protocol": "UDP",
      "src_ip": "10.80.130.5",
      "dst_ip": "10.80.150.11",
      "src_port": 514,
      "dst_port": 514,
      "action": "forward"
    }
  ],
  "custom_dynamic": {
    "unexpected_field_1": "random_value",
    "nested_map": {
      "innerKeyA": "A",
      "innerKeyB": 123,
      "innerKeyC": [true, false, "maybe"]
    },
    "array_of_objects": [
      {"key": "value1"},
      {"key": "value2", "extra": 999}
    ]
  },
  "timestamp": "2025-10-22T15:17:16.948+0700"
}
`
const AuthSuccessTestData = `
{
  "rule": {
    "groups": ["authentication_success", "windows"]
  },
  "data": {
    "dstuser": "testuser"
  }
}
`

const PrivilegeEscalationTestData = `
{
  "predecoder": {
    "program_name": "sudo"
  },
  "data": {
    "srcuser": "root",
    "dstuser": "admin"
  }
}
`

const PostgreSQLTestData = `
{
  "agent": {
    "ip": "10.80.100.17"
  },
  "location": "/var/lib/pgsql/15/data/log/postgresql-2025-10-22.log"
}
`

const SeverityTestData = `
{
  "rule": {
    "level": 8
  }
}
`
