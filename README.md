# goosquery

Go osquery with JSON output, currently working on Windows.

The goal is to provide a lightweight, portable, and easy-to-use version of osquery that can be integrated into other projects.

**Note**: Currently does not exclude records generated by the program itself.

## Status

| Status       | Icon | Description                                                                 |
|--------------|------|-----------------------------------------------------------------------------|
| Not Started  | ⏳   | Work has not begun on this table yet.                                       |
| In Progress  | 🛠️   | Actively being developed.                                                   |
| Completed    | ✅   | Fully implemented and tested.                                               |
| Testing      | 🧪   | Development is done, but under testing for bugs or issues.                  |
| Blocked      | ⛔   | Development is paused due to dependencies, blockers, or technical issues.   |
| Planned      | 🗓️   | Table is planned for future implementation but hasn’t started yet.          |
| Deprecated   | 🗑️   | This table is no longer relevant or supported in this implementation.       |

## Table

| Table Name                       | Status  |
|----------------------------------|---------|
| appcompat_shims                  | 🧪      |
| arp_cache                        | 🧪      |
| authenticode                     | 🛠️      |
| autoexec                         | ⏳      |
| azure_instance_metadata          | ⏳      |
| azure_instance_tags              | ⏳      |
| background_activities_moderator  | 🧪      |
| battery                          | ⏳      |
| bitlocker_info                   | 🧪      |
| carbon_black_info                | ⏳      |
| carves                           | ⏳      |
| certificates                     | 🛠️      |
| chassis_info                     | ✅      |
| chocolatey_packages              | ⏳      |
| chrome_extension_content_scripts | ⏳      |
| chrome_extensions                | ⏳      |
| connectivity                     | 🧪      |
| cpu_info                         | ✅      |
| cpuid                            | ⏳      |
| curl                             | ⏳      |
| curl_certificate                 | ⏳      |
| default_environment              | 🧪      |
| deviceguard_status               | ✅      |
| disk_info                        | ✅      |
| dns_cache                        | ✅      |
| drivers                          | 🧪      |
| ec2_instance_metadata            | ⏳      |
| ec2_instance_tags                | ⏳      |
| etc_hosts                        | ✅      |
| etc_protocols                    | ✅      |
| etc_services                     | ✅      |
| file                             | 🛠️      |
| firefox_addons                   | ⏳      |
| groups                           | ⏳      |
| hash                             | ⏳      |
| ie_extensions                    | ⏳      |
| intel_me_info                    | ⏳      |
| interface_addresses              | ✅      |
| interface_details                | 🧪      |
| kernel_info                      | ✅      |
| kva_speculative_info             | ✅      |
| listening_ports                  | ✅      |
| logged_in_users                  | 🛠️      |
| logical_drives                   | ✅      |
| logon_sessions                   | 🧪      |
| memory_devices                   | 🧪      |
| npm_packages                     | ⏳      |
| ntdomains                        | ⏳      |
| ntfs_acl_permissions             | ⏳      |
| ntfs_journal_events              | ⏳      |
| office_mru                       | ⏳      |
| os_version                       | 🧪      |
| osquery_events                   | 🗑️      |
| osquery_extensions               | 🗑️      |
| osquery_flags                    | 🗑️      |
| osquery_info                     | 🗑️      |
| osquery_packs                    | 🗑️      |
| osquery_registry                 | 🗑️      |
| osquery_schedule                 | 🗑️      |
| patches                          | ✅      |
| physical_disk_performance        | ⏳      |
| pipes                            | 🧪      |
| platform_info                    | 🧪      |
| powershell_events                | ⏳      |
| prefetch                         | ⏳      |
| process_etw_events               | ⏳      |
| process_memory_map               | ✅      |
| process_open_sockets             | ✅      |
| processes                        | 🧪      |
| programs                         | 🧪      |
| python_packages                  | 🧪      |
| registry                         | 🧪      |
| routes                           | 🛠️      |
| scheduled_tasks                  | 🧪      |
| secureboot                       | ⏳      |
| security_profile_info            | 🛠️      |
| services                         | ✅      |
| shared_resources                 | ✅      |
| shellbags                        | ⏳      |
| shimcache                        | ⏳      |
| ssh_configs                      | ⏳      |
| startup_items                    | 🛠️      |
| system_info                      | 🧪      |
| time                             | ✅      |
| tpm_info                         | ⏳      |
| uptime                           | ✅      |
| user_groups                      | 🛠️      |
| user_ssh_keys                    | ⏳      |
| userassist                       | ⏳      |
| users                            | ✅      |
| video_info                       | ⏳      |
| vscode_extensions                | ⏳      |
| winbaseobj                       | 🧪      |
| windows_crashes                  | ⏳      |
| windows_eventlog                 | ⏳      |
| windows_events                   | ⏳      |
| windows_firewall_rules           | 🧪      |
| windows_optional_features        | ✅      |
| windows_search                   | ⏳      |
| windows_security_center          | 🧪      |
| windows_security_products        | 🛠️      |
| windows_update_history           | 🛠️      |
| wmi_bios_info                    | ⏳      |
| wmi_cli_event_consumers          | ⏳      |
| wmi_event_filters                | ⏳      |
| wmi_filter_consumer_binding      | ⏳      |
| wmi_script_event_consumers       | ⏳      |
| yara                             | ⛔      |
| yara_events                      | ⛔      |
| ycloud_instance_metadata         | ⛔      |
