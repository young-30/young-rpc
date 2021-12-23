## 1-评分趋势接口
* 使用场景
> 获取某一时间段内各大应用市场对产品的评分趋势
> * 接口类型    
> `HTTP + JSON`
* HTTP Method 
> GET
* URL
> <http://IP:PORT/product/service/performance/get_score_trend>
* 请求参数

|   参数名    |  类型  | 必选 | 默认值 |                             说明                             |
| :---------: | :----: | :--: | :----: | :----------------------------------------------------------: |
|   uniq_id   |  int   | yes  |  NULL  |                       事件ID(带版本号)                       |
| start_time  | string | yes  |  NULL  |                           起始时间                           |
|  end_time   | string | yes  |  NULL  |                           结束时间                           |
| app_channel | String | yes  |  NULL  |                 目标应用市场的三级渠道ID列表                 |
|  filter_id  |  int   | yes  |  NULL  | 数据聚合方式，[{name:  '评分', id: 1001}, { name: '总量', id: 1002}] |
|   channel   | string |  no  |  NULL  |                   渠道，传入规则详见1.0.1                    |

* 返回参数

|    参数名     |  类型  |                          说明                          |
| :-----------: | :----: | :----------------------------------------------------: |
| result: date  | string |                当日日期(精确到年/月/日)                |
| result: score | float  | 产品在该应用市场下的下的当日平均得分 (filter_way=1001) |
|  result: num  |  int   | 产品在该应用市场下的下的当日评论数量 (filter_way=1002) |
| third_channel |  int   |                 该应用市场的三级渠道ID                 |

* 请求示例1 
> <http://127.0.0.1/product/service/performance/get_score_trend?uniq_id=2000000170003&start_time=2020-10-04&end_time=2020-10-06&app_channel=4,7,9&filter_id=1001>

* 返回示例1


        {
          "msg": "success",
          "code": 0,
          "res": [
            {
              "third_channel": 4,
              "result": [
                {
                  "date": "2020-10-04",
                  "value": 3.8
                },
                {
                  "date": "2020-10-05",
                  "value": 4.1
                },
                {
                  "date": "2020-10-06",
                  "value": 2.4
                }
              ],  
            },
            {
              "third_channel": 7,
              "result": [
                {
                  "date": "2020-10-04",
                  "value": 4.2
                },
                {
                  "date": "2020-10-05",
                  "value": 4.8
                },
                {
                  "date": "2020-10-06",
                  "value": 3.1
                }
              ],
            },
            ...
          ]
        }


* 请求示例2 
> <http://127.0.0.1/product/service/performance/get_score_num_trend?uniq_id=2000000170003&start_time=2020-10-04&end_time=2020-10-06&app_channel=4,7,9&filter_id=1002>

* 返回示例2


        {
          "msg": "success",
          "code": 0,
          "res": [
            {
              "third_channel": 4,
              "result": [
                {
                  "date": "2020-10-04",
                  "value": 2001
                },
                {
                  "date": "2020-10-05",
                  "value": 3002
                },
                {
                  "date": "2020-10-06",
                  "value": 4004
                }
              ],
            },
            {
              "third_channel": 7,
              "result": [
                {
                  "date": "2020-10-04",
                  "value": 1001
                },
                {
                  "date": "2020-10-05",
                  "value": 3000
                },
                {
                  "date": "2020-10-06",
                  "value": 5001
                }
              ],
            },
            ...
          ]
        }

***



## 2-评分计数接口

* 使用场景
> 获取某一时间段内具体应用市场对产品的评分详情
> * 接口类型    
> `HTTP + JSON`
* HTTP Method
> GET
* URL
> <http://IP:PORT/product/service/performance/get_score_detail>

* 请求参数

|   参数名    |  类型  | 必选 | 默认值 |             说明             |
| :---------: | :----: | :--: | :----: | :--------------------------: |
|   uniq_id   |  int   | yes  |  NULL  |       事件ID(带版本号)       |
| start_time  | string | yes  |  NULL  |           起始时间           |
|  end_time   | string | yes  |  NULL  |           结束时间           |
| app_channel | table  | yes  |  NULL  | 目标应用市场的三级渠道ID列表 |
|   channel   | string |  no  |  NULL  |   渠道，传入规则详见1.0.1    |

* 返回参数

|    参数名     | 类型  |            说明            |
| :-----------: | :---: | :------------------------: |
| third_channel |  int  |   该应用市场的三级渠道ID   |
| result.star_* |  int  |      对应星级的评分量      |
| result.total  |  int  | 该应用市场下的所有评论数量 |
|  result.avg   | float |   该应用市场下的平均得分   |

* 请求示例 
> <http://127.0.0.1/product/service/performance/get_score_detail?uniq_id=2000000170003&start_time=2020-10-04&end_time=2020-10-06&app_channel=4,7,9>

* 返回示例


        {
          "msg": "success",
          "code": 0,
          "res": [
            {
              "result": 
                {
                  “star_5": 1326,
                  "star_4": 206,
                  "star_3": 435,
                  "star_2": 258,
                  "star_1": 311,
                  "total": 2536,
                  "avg": 3.7
                },
              "third_channel": 4,
            },
            {
              "result": 
                {
                  "star_5": 1400,
                  "star_4": 100,
                  "star_3": 200,
                  "star_2": 300,
                  "star_1": 100,
                  "total": 2100,
                  "avg": 4.1
                },
              "third_channel": 7,
            },
            ...
          ]
        }
***



## 3-日/周/月评分接口   

* 使用场景
> 获取某一时间段内具体应用市场对产品的评分日期分布
> * 接口类型    
> `HTTP + JSON`
* HTTP Method
> GET
* URL

> <http://IP:PORT/product/service/performance/get_score_by_date>
* 请求参数

|   参数名    |  类型  | 必选 | 默认值 |                         说明                         |
| :---------: | :----: | :--: | :----: | :--------------------------------------------------: |
|   uniq_id   |  int   | yes  |  NULL  |                   事件ID(带版本号)                   |
| start_time  | string | yes  |  NULL  |                       起始时间                       |
|  end_time   | string | yes  |  NULL  |                       结束时间                       |
| app_channel | table  | yes  |  NULL  |             目标应用市场的三级渠道ID列表             |
|  date_unit  | string | yes  |  NULL  | 聚合数据的日期单位，可选值包括"day"、"week"、"month" |
|   channel   | string |  no  |  NULL  |               渠道，传入规则详见1.0.1                |

* 返回参数

|     参数名     |  类型  |                  说明                  |
| :------------: | :----: | :------------------------------------: |
| third_channel  |  int   |         该应用市场的三级渠道ID         |
|   date_unit    | string |           聚合数据的日期单位           |
| result: total  |  int   |              所有评分数量              |
|  result: date  | string | 当前日期(按周聚合时为对应周周一的日期) |
| result: star_* |  int   |           对应星级的评分数量           |

* 请求示例1
> <http://127.0.0.1/product/service/performance/get_score_by_date?uniq_id=2000000170003&start_time=2020-10-04&end_time=2020-10-06&app_channel=23,25,27&date_unit=day>

* 返回示例1


    {
      "msg": "success",
      "code": 0,
      "res": [
        {
          "third_channel": 23,
          "date_unit": "day",
          "result": [
            {
              "total": 5977,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-04"
            },
            {
              "total": 5808,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-05"
            },
            {
              "total": 5924,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-06"
            }
          ]
        },
        {
          "third_channel": 25,
          "date_unit": "day",
          "result": [
            {
              "total": 3823,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-04"
            },
            {
              "total": 3925,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-05"
            },
            {
              "total": 3730,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-06"
            }
          ]
        },
        {
          "third_channel": 27,
          "date_unit": "day",
          "result": [
            {
              "total": 695,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-04"
            },
            {
              "total": 716,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-05"
            },
            {
              "total": 686,
              "star_1": 0,
              "star_2": 0,
              "star_3": 0,
              "star_4": 0,
              "star_5": 0,
              "date": "2020-10-06"
            }
          ]
        }
      ]
    }



* 请求示例2
> <http://127.0.0.1/product/service/performance/get_score_by_date?uniq_id=2000000170003&start_time=2020-09-01&end_time=2020-09-19&app_channel=4,7,9&date_unit=week>

* 返回示例2
```
{
  "msg": "success",
  "code": 0,
  "res": [
    {
      "third_channel": 23,
      "date_unit": "week",
      "result": [
        {
          "total": 34663,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-08-31"
        },
        {
          "total": 29619,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-07"
        },
        {
          "total": 38538,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-14"
        }
      ]
    },
    {
      "third_channel": 25,
      "date_unit": "week",
      "result": [
        {
          "total": 22529,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-08-31"
        },
        {
          "total": 22664,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-07"
        },
        {
          "total": 27644,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-14"
        }
      ]
    },
    {
      "third_channel": 27,
      "date_unit": "week",
      "result": [
        {
          "total": 3000,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-08-31"
        },
        {
          "total": 2762,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-07"
        },
        {
          "total": 3464,
          "star_1": 0,
          "star_2": 0,
          "star_3": 0,
          "star_4": 0,
          "star_5": 0,
          "date": "2020-09-14"
        }
      ]
    }
  ]
}
```

  

​    




