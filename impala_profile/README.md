## impala_profile

> impala_profile processor 解析Impala Profile日志，放入 event 中。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - impala_profile:
      # 源字段，impala_profile processor 从此字段读取日志
      # 此配置默认值为 "message"
      field: "message"
      # 目标字段，impala_profile processor 解析成功后将其写入到目标字段中
      # 此配置默认值为"impala_profile"
      target_field:"impala_profile"
      # 自定义固定映射，impala_profile processor 解析成功后，用里面的Key/Value替换原有值
      # 此配置默认值
      const_mappings: 
			domain:     "Impala",
			logLevel:   "INFO",
			eventName:  "Profile",
			threadName: "MAIN",
	  # processor 标记位，impala_profile processor 处理成功后会将此字段设置为 true
      # 通常该字段用于标识作用，方便后面的 logstash 判断 event 是否被某个 processor 处理过
      # 此配置默认值为 "processors.impala_profile"
      processors_field: "processors.impala_profile"
      # 当发现相同字段时，如果该配置为 true，则覆盖event已有的值
      # 此配置默认值为 true
      overwrite_keys: true
      # 当无法找到 source_field 指定的字段时，如果该配置为 true，则忽略错误，继续处理 event
      # 此配置默认值为 false
      ignore_missing: true
      # 当出现一些错误时(例如上面的 source_field 找不到或者 source_field 不是个字符串等)忽略
      # 错误继续处理 event，可以将 ignore_failure 视为 ignore_missing 的更大范畴兼容
      # 此配置默认值为 true
      ignore_failure: true
```

### 如何调试?

你可以为 logstash 开启终端输出来实时观察日志处理情况:

``` sh
output {
  stdout {
    codec => rubydebug
  }
}
```

如果 impala_profile processor 处理成功后应该可以在 logstash 控制台看到 `target_fileds` 字段

``` diff
  ],
  "input": {
    "type": "filestream"
  },
  "ecs": {
    "version": "8.0.0"
  },
  "host": {
    "name": "XiangLi-T14"
  },
  "agent": {
    "type": "filebeat",
    "version": "8.3.3",
    "ephemeral_id": "2d1524b1-a4f4-4947-bd64-c75215d66eee",
    "id": "59d8e91f-4173-48a2-8ff3-3f8d61f07891",
    "name": "XiangLi-T14"
  },
+  "impala_profile": {
+    "eventName": "Profile",
+    "threadName": "MAIN",
+    "domain": "Impala",
+    "timestamp": "2022-08-19T11:35:47+08:00",
+    "query_id": "804943a398291a99:9b9353ee00000000",
+    "profile": "TRuntimeProfileTree({Nodes:[TRuntimeProfileNode({Name:Query (id=804943a398291a99:9b9353ee00000000) NumChildren:2 Counters:[TCounter({Name:InactiveTotalTime Unit:TIME_NS Value:0}) TCounter({Name:TotalTime Unit:TIME_NS Value:0})] Metadata:-1 Indent:true InfoStrings:map[] InfoStringsDisplayOrder:[] ChildCountersMap:map[] EventSequences:[] TimeSeriesCounters:[] SummaryStatsCounters:[] NodeMetadata:<nil> Aggregated:<nil>}) TRuntimeProfileNode({Name:Summary NumChildren:0 Counters:[TCounter({Name:InactiveTotalTime Unit:TIME_NS Value:0}) TCounter({Name:TotalTime Unit:TIME_NS Value:0})] Metadata:-1 Indent:true InfoStrings:map[Connected User:apprc@PANEL.COM Coordinator:e4bdhpp023:22000 Default Db:dw01 Delegated User:zouyue01 End Time:2022-08-19 11:35:47.430362000 HiveServer2 Protocol Version:V6 Impala Version:impalad version 2.12.0-RELEASE RELEASE (build 56d0ad4872bc05d499438c75e05bb6c85c3a3b83) Network Address:::ffff:11.11.237.129:55588 Query Options (set by configuration and planner):MEM_LIMIT=8589934592 Query Options (set by configuration):MEM_LIMIT=8589934592 Query State:EXCEPTION Query Status:Invalid query option: useSSL\n Query Type:SET Session ID:ef46d78fc019b3a0:3c6515810c874eb4 Session Type:HIVESERVER2 Sql Statement:SET useSSL=false Start Time:2022-08-19 11:35:47.348288000 User:zouyue01] InfoStringsDisplayOrder:[Session ID Session Type HiveServer2 Protocol Version Start Time End Time Query Type Query State Query Status Impala Version User Connected User Delegated User Network Address Default Db Sql Statement Coordinator Query Options (set by configuration) Query Options (set by configuration and planner)] ChildCountersMap:map[] EventSequences:[TEventSequence({Name: Timestamps:[] Labels:[]}) TEventSequence({Name:Query Timeline Timestamps:[89159 866704 82061087 82073607] Labels:[Query submitted Planning finished Cancelled Unregister query]})] TimeSeriesCounters:[] SummaryStatsCounters:[] NodeMetadata:<nil> Aggregated:<nil>}) TRuntimeProfileNode({Name:ImpalaServer NumChildren:0 Counters:[TCounter({Name:ClientFetchWaitTimer Unit:TIME_NS Value:0}) TCounter({Name:InactiveTotalTime Unit:TIME_NS Value:0}) TCounter({Name:RowMaterializationTimer Unit:TIME_NS Value:0}) TCounter({Name:TotalTime Unit:TIME_NS Value:0})] Metadata:-1 Indent:true InfoStrings:map[] InfoStringsDisplayOrder:[] ChildCountersMap:map[:[ClientFetchWaitTimer RowMaterializationTimer]] EventSequences:[] TimeSeriesCounters:[] SummaryStatsCounters:[] NodeMetadata:<nil> Aggregated:<nil>})] ExecSummary:<nil> ProfileVersion:<nil>})",
+    "logLevel": "INFO"
+  },
  "log": {
    "offset": 427540,
    "file": {
      "path": "/home/xiangli/go_dev/go_root/src/github.com/elastic/beats/filebeat/tests/files/logs/impala_profile_log"
    }
  },
  "message": "1660880147430 804943a398291a99:9b9353ee00000000 eJyVVM1O20AQNi1qQ6D8CWhSirSVegAJovVfYltFKkqMiJQEikPorVp7J7CqY4e1DQpPUPXAE/TQ9kH6LD313luP3cTQJhFUdGXJ2tmd+eab+Xbyr3KbbxLgPbTO6LaBNVNTiWoaiikT07RM11R1FQBfr43lyfxmbqEaEC9m59AMY+I3WQeWsyuSlJsa3a9MLKxK+cyqOHnsJJ0O4b1l6f/cFz/MlsMgAC8Gio4i4HOk2+Xe64Odhl0rlPfr0+Uw5JQFJA75PGguPe12saJaiiKyzVagTRI/RhV3kl5gebYCPpyQm1iZyzDpJYDljB1Q1MddU7CibGFjSzaRLFuqbmmlgqZitdgP93xP5OwAPweuoAMexqEX+qgFPGJh8KBVnK12usQnN5ZjNthSdJ7ukVKQlQLeOrRr9o5jo5v/upswnyK9SDGhmlFSXA/rVDNFJwyvpAPWXbfoGbonGuMa6sZcA+KLkL9HO5RyiKJnltUWy5LlgvgUtSRgTEvXdcPAaWv3u7HAj9B6BDFye8gLgzY7STjpmxER5Ls+EVXmG0t1u/6uVq1Xm9uGbpimqumm8vIeUW73nE49nVjUfMp+W7YPmtX9xsxfaxKtVYNz4jOKzgbGcABioSQCx6lls+nVZq8LDx27mXUE337O1coLaGtFWjLaHpZNVyXYUr2iLuuGjD2jpIGrzdxc7ntP71VbtmMftuxD5Ylz5qc5dSCI50Xca7jtNvEjyIojHt+tB1UzFFFajCdHRJT/tTiU3gj4P4UzBPdHh0O0h0s4UrgxtQ2SGXssY3ofF87Q8xgtyfCjulfz/1tnq/05IOUf5TNSbvaarODtswDyu1c/s59/sG+fvjSuvn9t5Pfm0gtR4nZYLMgsHPSjsOAEtVnAolOgU2USeOD7QOePAg4nLIqBp4oSc2UmrVTaADGBarmlss8Ez12IvdNjwuI+Nk+H0F3D6elheFEX9eFMqPVywGjI647JN/FRWr8V665okvQbBTnANg=="
}
```
