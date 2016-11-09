package holidayJP

import (
    "encoding/json"
    "fmt"
    "time"
)

var _holidays map[string]string

func holidayJson() string {
    text := `
 {
  "2016-01-01": "元日",
  "2016-01-11": "成人の日",
  "2016-02-11": "建国記念の日",
  "2016-03-20": "春分の日",
  "2016-03-21": "振替休日",
  "2016-04-29": "昭和の日",
  "2016-05-03": "憲法記念日",
  "2016-05-04": "みどりの日",
  "2016-05-05": "こどもの日",
  "2016-07-18": "海の日",
  "2016-08-11": "山の日",
  "2016-09-19": "敬老の日",
  "2016-09-22": "秋分の日",
  "2016-10-10": "体育の日",
  "2016-11-03": "文化の日",
  "2016-11-23": "勤労感謝の日",
  "2016-12-23": "天皇誕生日",
  "2017-01-01": "元日",
  "2017-01-02": "振替休日",
  "2017-01-09": "成人の日",
  "2017-02-11": "建国記念の日",
  "2017-02-23": "天皇誕生日",
  "2017-03-20": "春分の日",
  "2017-04-29": "昭和の日",
  "2017-05-03": "憲法記念日",
  "2017-05-04": "みどりの日",
  "2017-05-05": "こどもの日",
  "2017-07-17": "海の日",
  "2017-08-11": "山の日",
  "2017-09-18": "敬老の日",
  "2017-09-23": "秋分の日",
  "2017-10-09": "体育の日",
  "2017-11-03": "文化の日",
  "2017-11-23": "勤労感謝の日",
  "2018-01-01": "元日",
  "2018-01-08": "成人の日",
  "2018-02-11": "建国記念の日",
  "2018-02-12": "振替休日",
  "2018-02-23": "天皇誕生日",
  "2018-03-21": "春分の日",
  "2018-04-29": "昭和の日",
  "2018-04-30": "振替休日",
  "2018-05-03": "憲法記念日",
  "2018-05-04": "みどりの日",
  "2018-05-05": "こどもの日",
  "2018-07-16": "海の日",
  "2018-08-11": "山の日",
  "2018-09-17": "敬老の日",
  "2018-09-23": "秋分の日",
  "2018-09-24": "振替休日",
  "2018-10-08": "体育の日",
  "2018-11-03": "文化の日",
  "2018-11-23": "勤労感謝の日",
  "2018-12-24": "振替休日",
  "2019-01-01": "元日",
  "2019-01-14": "成人の日",
  "2019-02-11": "建国記念の日",
  "2019-02-23": "天皇誕生日",
  "2019-03-21": "春分の日",
  "2019-04-29": "昭和の日",
  "2019-05-03": "憲法記念日",
  "2019-05-04": "みどりの日",
  "2019-05-05": "こどもの日",
  "2019-05-06": "振替休日",
  "2019-07-15": "海の日",
  "2019-08-11": "山の日",
  "2019-08-12": "振替休日",
  "2019-09-16": "敬老の日",
  "2019-09-23": "秋分の日",
  "2019-10-14": "体育の日",
  "2019-11-03": "文化の日",
  "2019-11-04": "振替休日",
  "2019-11-23": "勤労感謝の日",
  "2020-01-01": "元日",
  "2020-01-13": "成人の日",
  "2020-02-11": "建国記念の日",
  "2020-02-23": "天皇誕生日",
  "2020-03-20": "春分の日",
  "2020-04-29": "昭和の日",
  "2020-05-03": "憲法記念日",
  "2020-05-04": "みどりの日",
  "2020-05-05": "こどもの日",
  "2020-05-06": "振替休日",
  "2020-07-20": "海の日",
  "2020-08-11": "山の日",
  "2020-09-21": "敬老の日",
  "2020-09-22": "秋分の日",
  "2020-10-12": "体育の日",
  "2020-11-03": "文化の日",
  "2020-11-23": "勤労感謝の日",
  "2021-01-01": "元日",
  "2021-01-11": "成人の日",
  "2021-02-11": "建国記念の日",
  "2021-02-23": "天皇誕生日",
  "2021-03-20": "春分の日",
  "2021-04-29": "昭和の日",
  "2021-05-03": "憲法記念日",
  "2021-05-04": "みどりの日",
  "2021-05-05": "こどもの日",
  "2021-07-19": "海の日",
  "2021-08-11": "山の日",
  "2021-09-20": "敬老の日",
  "2021-09-23": "秋分の日",
  "2021-10-11": "体育の日",
  "2021-11-03": "文化の日",
  "2021-11-23": "勤労感謝の日",
  "2022-01-01": "元日",
  "2022-01-10": "成人の日",
  "2022-02-11": "建国記念の日",
  "2022-02-23": "天皇誕生日",
  "2022-03-21": "春分の日",
  "2022-04-29": "昭和の日",
  "2022-05-03": "憲法記念日",
  "2022-05-04": "みどりの日",
  "2022-05-05": "こどもの日",
  "2022-07-18": "海の日",
  "2022-08-11": "山の日",
  "2022-09-19": "敬老の日",
  "2022-09-23": "秋分の日",
  "2022-10-10": "体育の日",
  "2022-11-03": "文化の日",
  "2022-11-23": "勤労感謝の日",
  "2023-01-01": "元日",
  "2023-01-02": "振替休日",
  "2023-01-09": "成人の日",
  "2023-02-11": "建国記念の日",
  "2023-02-23": "天皇誕生日",
  "2023-03-21": "春分の日",
  "2023-04-29": "昭和の日",
  "2023-05-03": "憲法記念日",
  "2023-05-04": "みどりの日",
  "2023-05-05": "こどもの日",
  "2023-07-17": "海の日",
  "2023-08-11": "山の日",
  "2023-09-18": "敬老の日",
  "2023-09-23": "秋分の日",
  "2023-10-09": "体育の日",
  "2023-11-03": "文化の日",
  "2023-11-23": "勤労感謝の日",
  "2024-01-01": "元日",
  "2024-01-08": "成人の日",
  "2024-02-11": "建国記念の日",
  "2024-02-12": "振替休日",
  "2024-02-23": "天皇誕生日",
  "2024-03-20": "春分の日",
  "2024-04-29": "昭和の日",
  "2024-05-03": "憲法記念日",
  "2024-05-04": "みどりの日",
  "2024-05-05": "こどもの日",
  "2024-05-06": "振替休日",
  "2024-07-15": "海の日",
  "2024-08-11": "山の日",
  "2024-08-12": "振替休日",
  "2024-09-16": "敬老の日",
  "2024-09-22": "秋分の日",
  "2024-09-23": "振替休日",
  "2024-10-14": "体育の日",
  "2024-11-03": "文化の日",
  "2024-11-04": "振替休日",
  "2024-11-23": "勤労感謝の日",
  "2025-01-01": "元日",
  "2025-01-13": "成人の日",
  "2025-02-11": "建国記念の日",
  "2025-02-23": "天皇誕生日",
  "2025-03-20": "春分の日",
  "2025-04-29": "昭和の日",
  "2025-05-03": "憲法記念日",
  "2025-05-04": "みどりの日",
  "2025-05-05": "こどもの日",
  "2025-05-06": "振替休日",
  "2025-07-21": "海の日",
  "2025-08-11": "山の日",
  "2025-09-15": "敬老の日",
  "2025-09-23": "秋分の日",
  "2025-10-13": "体育の日",
  "2025-11-03": "文化の日",
  "2025-11-23": "勤労感謝の日",
  "2025-11-24": "振替休日",
  "2026-01-01": "元日",
  "2026-01-12": "成人の日",
  "2026-02-11": "建国記念の日",
  "2026-02-23": "天皇誕生日",
  "2026-03-20": "春分の日",
  "2026-04-29": "昭和の日",
  "2026-05-03": "憲法記念日",
  "2026-05-04": "みどりの日",
  "2026-05-05": "こどもの日",
  "2026-05-06": "振替休日",
  "2026-07-20": "海の日",
  "2026-08-11": "山の日",
  "2026-09-21": "敬老の日",
  "2026-09-22": "国民の休日",
  "2026-09-23": "秋分の日",
  "2026-10-12": "体育の日",
  "2026-11-03": "文化の日",
  "2026-11-23": "勤労感謝の日",
  "2027-01-01": "元日",
  "2027-01-11": "成人の日",
  "2027-02-11": "建国記念の日",
  "2027-02-23": "天皇誕生日",
  "2027-03-21": "春分の日",
  "2027-03-22": "振替休日",
  "2027-04-29": "昭和の日",
  "2027-05-03": "憲法記念日",
  "2027-05-04": "みどりの日",
  "2027-05-05": "こどもの日",
  "2027-07-19": "海の日",
  "2027-08-11": "山の日",
  "2027-09-20": "敬老の日",
  "2027-09-23": "秋分の日",
  "2027-10-11": "体育の日",
  "2027-11-03": "文化の日",
  "2027-11-23": "勤労感謝の日",
  "2028-01-01": "元日",
  "2028-01-10": "成人の日",
  "2028-02-11": "建国記念の日",
  "2028-02-23": "天皇誕生日",
  "2028-03-20": "春分の日",
  "2028-04-29": "昭和の日",
  "2028-05-03": "憲法記念日",
  "2028-05-04": "みどりの日",
  "2028-05-05": "こどもの日",
  "2028-07-17": "海の日",
  "2028-08-11": "山の日",
  "2028-09-18": "敬老の日",
  "2028-09-22": "秋分の日",
  "2028-10-09": "体育の日",
  "2028-11-03": "文化の日",
  "2028-11-23": "勤労感謝の日",
  "2029-01-01": "元日",
  "2029-01-08": "成人の日",
  "2029-02-11": "建国記念の日",
  "2029-02-12": "振替休日",
  "2029-02-23": "天皇誕生日",
  "2029-03-20": "春分の日",
  "2029-04-29": "昭和の日",
  "2029-04-30": "振替休日",
  "2029-05-03": "憲法記念日",
  "2029-05-04": "みどりの日",
  "2029-05-05": "こどもの日",
  "2029-07-16": "海の日",
  "2029-08-11": "山の日",
  "2029-09-17": "敬老の日",
  "2029-09-23": "秋分の日",
  "2029-09-24": "振替休日",
  "2029-10-08": "体育の日",
  "2029-11-03": "文化の日",
  "2029-11-23": "勤労感謝の日",
  "2029-12-24": "振替休日",
  "2030-01-01": "元日",
  "2030-01-14": "成人の日",
  "2030-02-11": "建国記念の日",
  "2030-02-23": "天皇誕生日",
  "2030-03-20": "春分の日",
  "2030-04-29": "昭和の日",
  "2030-05-03": "憲法記念日",
  "2030-05-04": "みどりの日",
  "2030-05-05": "こどもの日",
  "2030-05-06": "振替休日",
  "2030-07-15": "海の日",
  "2030-08-11": "山の日",
  "2030-08-12": "振替休日",
  "2030-09-16": "敬老の日",
  "2030-09-23": "秋分の日",
  "2030-10-14": "体育の日",
  "2030-11-03": "文化の日",
  "2030-11-04": "振替休日",
  "2030-11-23": "勤労感謝の日",
  "2031-01-01": "元日",
  "2031-01-13": "成人の日",
  "2031-02-11": "建国記念の日",
  "2031-02-23": "天皇誕生日",
  "2031-03-21": "春分の日",
  "2031-04-29": "昭和の日",
  "2031-05-03": "憲法記念日",
  "2031-05-04": "みどりの日",
  "2031-05-05": "こどもの日",
  "2031-05-06": "振替休日",
  "2031-07-21": "海の日",
  "2031-08-11": "山の日",
  "2031-09-15": "敬老の日",
  "2031-09-23": "秋分の日",
  "2031-10-13": "体育の日",
  "2031-11-03": "文化の日",
  "2031-11-23": "勤労感謝の日",
  "2031-11-24": "振替休日",
  "2032-01-01": "元日",
  "2032-01-12": "成人の日",
  "2032-02-11": "建国記念の日",
  "2032-02-23": "天皇誕生日",
  "2032-03-20": "春分の日",
  "2032-04-29": "昭和の日",
  "2032-05-03": "憲法記念日",
  "2032-05-04": "みどりの日",
  "2032-05-05": "こどもの日",
  "2032-07-19": "海の日",
  "2032-08-11": "山の日",
  "2032-09-20": "敬老の日",
  "2032-09-21": "国民の休日",
  "2032-09-22": "秋分の日",
  "2032-10-11": "体育の日",
  "2032-11-03": "文化の日",
  "2032-11-23": "勤労感謝の日",
  "2033-01-01": "元日",
  "2033-01-10": "成人の日",
  "2033-02-11": "建国記念の日",
  "2033-02-23": "天皇誕生日",
  "2033-03-20": "春分の日",
  "2033-03-21": "振替休日",
  "2033-04-29": "昭和の日",
  "2033-05-03": "憲法記念日",
  "2033-05-04": "みどりの日",
  "2033-05-05": "こどもの日",
  "2033-07-18": "海の日",
  "2033-08-11": "山の日",
  "2033-09-19": "敬老の日",
  "2033-09-23": "秋分の日",
  "2033-10-10": "体育の日",
  "2033-11-03": "文化の日",
  "2033-11-23": "勤労感謝の日",
  "2034-01-01": "元日",
  "2034-01-02": "振替休日",
  "2034-01-09": "成人の日",
  "2034-02-11": "建国記念の日",
  "2034-02-23": "天皇誕生日",
  "2034-03-20": "春分の日",
  "2034-04-29": "昭和の日",
  "2034-05-03": "憲法記念日",
  "2034-05-04": "みどりの日",
  "2034-05-05": "こどもの日",
  "2034-07-17": "海の日",
  "2034-08-11": "山の日",
  "2034-09-18": "敬老の日",
  "2034-09-23": "秋分の日",
  "2034-10-09": "体育の日",
  "2034-11-03": "文化の日",
  "2034-11-23": "勤労感謝の日",
  "2035-01-01": "元日",
  "2035-01-08": "成人の日",
  "2035-02-11": "建国記念の日",
  "2035-02-12": "振替休日",
  "2035-02-23": "天皇誕生日",
  "2035-03-21": "春分の日",
  "2035-04-29": "昭和の日",
  "2035-04-30": "振替休日",
  "2035-05-03": "憲法記念日",
  "2035-05-04": "みどりの日",
  "2035-05-05": "こどもの日",
  "2035-07-16": "海の日",
  "2035-08-11": "山の日",
  "2035-09-17": "敬老の日",
  "2035-09-23": "秋分の日",
  "2035-09-24": "振替休日",
  "2035-10-08": "体育の日",
  "2035-11-03": "文化の日",
  "2035-11-23": "勤労感謝の日",
  "2035-12-24": "振替休日",
  "2036-01-01": "元日",
  "2036-01-14": "成人の日",
  "2036-02-11": "建国記念の日",
  "2036-02-23": "天皇誕生日",
  "2036-03-20": "春分の日",
  "2036-04-29": "昭和の日",
  "2036-05-03": "憲法記念日",
  "2036-05-04": "みどりの日",
  "2036-05-05": "こどもの日",
  "2036-05-06": "振替休日",
  "2036-07-21": "海の日",
  "2036-08-11": "山の日",
  "2036-09-15": "敬老の日",
  "2036-09-22": "秋分の日",
  "2036-10-13": "体育の日",
  "2036-11-03": "文化の日",
  "2036-11-23": "勤労感謝の日",
  "2036-11-24": "振替休日",
  "2037-01-01": "元日",
  "2037-01-12": "成人の日",
  "2037-02-11": "建国記念の日",
  "2037-02-23": "天皇誕生日",
  "2037-03-20": "春分の日",
  "2037-04-29": "昭和の日",
  "2037-05-03": "憲法記念日",
  "2037-05-04": "みどりの日",
  "2037-05-05": "こどもの日",
  "2037-05-06": "振替休日",
  "2037-07-20": "海の日",
  "2037-08-11": "山の日",
  "2037-09-21": "敬老の日",
  "2037-09-22": "国民の休日",
  "2037-09-23": "秋分の日",
  "2037-10-12": "体育の日",
  "2037-11-03": "文化の日",
  "2037-11-23": "勤労感謝の日",
  "2038-01-01": "元日",
  "2038-01-11": "成人の日",
  "2038-02-11": "建国記念の日",
  "2038-02-23": "天皇誕生日",
  "2038-03-20": "春分の日",
  "2038-04-29": "昭和の日",
  "2038-05-03": "憲法記念日",
  "2038-05-04": "みどりの日",
  "2038-05-05": "こどもの日",
  "2038-07-19": "海の日",
  "2038-08-11": "山の日",
  "2038-09-20": "敬老の日",
  "2038-09-23": "秋分の日",
  "2038-10-11": "体育の日",
  "2038-11-03": "文化の日",
  "2038-11-23": "勤労感謝の日",
  "2039-01-01": "元日",
  "2039-01-10": "成人の日",
  "2039-02-11": "建国記念の日",
  "2039-02-23": "天皇誕生日",
  "2039-03-21": "春分の日",
  "2039-04-29": "昭和の日",
  "2039-05-03": "憲法記念日",
  "2039-05-04": "みどりの日",
  "2039-05-05": "こどもの日",
  "2039-07-18": "海の日",
  "2039-08-11": "山の日",
  "2039-09-19": "敬老の日",
  "2039-09-23": "秋分の日",
  "2039-10-10": "体育の日",
  "2039-11-03": "文化の日",
  "2039-11-23": "勤労感謝の日",
  "2040-01-01": "元日",
  "2040-01-02": "振替休日",
  "2040-01-09": "成人の日",
  "2040-02-11": "建国記念の日",
  "2040-02-23": "天皇誕生日",
  "2040-03-20": "春分の日",
  "2040-04-29": "昭和の日",
  "2040-04-30": "振替休日",
  "2040-05-03": "憲法記念日",
  "2040-05-04": "みどりの日",
  "2040-05-05": "こどもの日",
  "2040-07-16": "海の日",
  "2040-08-11": "山の日",
  "2040-09-17": "敬老の日",
  "2040-09-22": "秋分の日",
  "2040-10-08": "体育の日",
  "2040-11-03": "文化の日",
  "2040-11-23": "勤労感謝の日",
  "2040-12-24": "振替休日",
  "2041-01-01": "元日",
  "2041-01-14": "成人の日",
  "2041-02-11": "建国記念の日",
  "2041-02-23": "天皇誕生日",
  "2041-03-20": "春分の日",
  "2041-04-29": "昭和の日",
  "2041-05-03": "憲法記念日",
  "2041-05-04": "みどりの日",
  "2041-05-05": "こどもの日",
  "2041-05-06": "振替休日",
  "2041-07-15": "海の日",
  "2041-08-11": "山の日",
  "2041-08-12": "振替休日",
  "2041-09-16": "敬老の日",
  "2041-09-23": "秋分の日",
  "2041-10-14": "体育の日",
  "2041-11-03": "文化の日",
  "2041-11-04": "振替休日",
  "2041-11-23": "勤労感謝の日",
  "2042-01-01": "元日",
  "2042-01-13": "成人の日",
  "2042-02-11": "建国記念の日",
  "2042-02-23": "天皇誕生日",
  "2042-03-20": "春分の日",
  "2042-04-29": "昭和の日",
  "2042-05-03": "憲法記念日",
  "2042-05-04": "みどりの日",
  "2042-05-05": "こどもの日",
  "2042-05-06": "振替休日",
  "2042-07-21": "海の日",
  "2042-08-11": "山の日",
  "2042-09-15": "敬老の日",
  "2042-09-23": "秋分の日",
  "2042-10-13": "体育の日",
  "2042-11-03": "文化の日",
  "2042-11-23": "勤労感謝の日",
  "2042-11-24": "振替休日",
  "2043-01-01": "元日",
  "2043-01-12": "成人の日",
  "2043-02-11": "建国記念の日",
  "2043-02-23": "天皇誕生日",
  "2043-03-21": "春分の日",
  "2043-04-29": "昭和の日",
  "2043-05-03": "憲法記念日",
  "2043-05-04": "みどりの日",
  "2043-05-05": "こどもの日",
  "2043-05-06": "振替休日",
  "2043-07-20": "海の日",
  "2043-08-11": "山の日",
  "2043-09-21": "敬老の日",
  "2043-09-22": "国民の休日",
  "2043-09-23": "秋分の日",
  "2043-10-12": "体育の日",
  "2043-11-03": "文化の日",
  "2043-11-23": "勤労感謝の日",
  "2044-01-01": "元日",
  "2044-01-11": "成人の日",
  "2044-02-11": "建国記念の日",
  "2044-02-23": "天皇誕生日",
  "2044-03-20": "春分の日",
  "2044-03-21": "振替休日",
  "2044-04-29": "昭和の日",
  "2044-05-03": "憲法記念日",
  "2044-05-04": "みどりの日",
  "2044-05-05": "こどもの日",
  "2044-07-18": "海の日",
  "2044-08-11": "山の日",
  "2044-09-19": "敬老の日",
  "2044-09-22": "秋分の日",
  "2044-10-10": "体育の日",
  "2044-11-03": "文化の日",
  "2044-11-23": "勤労感謝の日",
  "2045-01-01": "元日",
  "2045-01-02": "振替休日",
  "2045-01-09": "成人の日",
  "2045-02-11": "建国記念の日",
  "2045-02-23": "天皇誕生日",
  "2045-03-20": "春分の日",
  "2045-04-29": "昭和の日",
  "2045-05-03": "憲法記念日",
  "2045-05-04": "みどりの日",
  "2045-05-05": "こどもの日",
  "2045-07-17": "海の日",
  "2045-08-11": "山の日",
  "2045-09-18": "敬老の日",
  "2045-09-22": "秋分の日",
  "2045-10-09": "体育の日",
  "2045-11-03": "文化の日",
  "2045-11-23": "勤労感謝の日",
  "2046-01-01": "元日",
  "2046-01-08": "成人の日",
  "2046-02-11": "建国記念の日",
  "2046-02-12": "振替休日",
  "2046-02-23": "天皇誕生日",
  "2046-03-20": "春分の日",
  "2046-04-29": "昭和の日",
  "2046-04-30": "振替休日",
  "2046-05-03": "憲法記念日",
  "2046-05-04": "みどりの日",
  "2046-05-05": "こどもの日",
  "2046-07-16": "海の日",
  "2046-08-11": "山の日",
  "2046-09-17": "敬老の日",
  "2046-09-23": "秋分の日",
  "2046-09-24": "振替休日",
  "2046-10-08": "体育の日",
  "2046-11-03": "文化の日",
  "2046-11-23": "勤労感謝の日",
  "2046-12-24": "振替休日",
  "2047-01-01": "元日",
  "2047-01-14": "成人の日",
  "2047-02-11": "建国記念の日",
  "2047-02-23": "天皇誕生日",
  "2047-03-21": "春分の日",
  "2047-04-29": "昭和の日",
  "2047-05-03": "憲法記念日",
  "2047-05-04": "みどりの日",
  "2047-05-05": "こどもの日",
  "2047-05-06": "振替休日",
  "2047-07-15": "海の日",
  "2047-08-11": "山の日",
  "2047-08-12": "振替休日",
  "2047-09-16": "敬老の日",
  "2047-09-23": "秋分の日",
  "2047-10-14": "体育の日",
  "2047-11-03": "文化の日",
  "2047-11-04": "振替休日",
  "2047-11-23": "勤労感謝の日",
  "2048-01-01": "元日",
  "2048-01-13": "成人の日",
  "2048-02-11": "建国記念の日",
  "2048-02-23": "天皇誕生日",
  "2048-03-20": "春分の日",
  "2048-04-29": "昭和の日",
  "2048-05-03": "憲法記念日",
  "2048-05-04": "みどりの日",
  "2048-05-05": "こどもの日",
  "2048-05-06": "振替休日",
  "2048-07-20": "海の日",
  "2048-08-11": "山の日",
  "2048-09-21": "敬老の日",
  "2048-09-22": "秋分の日",
  "2048-10-12": "体育の日",
  "2048-11-03": "文化の日",
  "2048-11-23": "勤労感謝の日",
  "2049-01-01": "元日",
  "2049-01-11": "成人の日",
  "2049-02-11": "建国記念の日",
  "2049-02-23": "天皇誕生日",
  "2049-03-20": "春分の日",
  "2049-04-29": "昭和の日",
  "2049-05-03": "憲法記念日",
  "2049-05-04": "みどりの日",
  "2049-05-05": "こどもの日",
  "2049-07-19": "海の日",
  "2049-08-11": "山の日",
  "2049-09-20": "敬老の日",
  "2049-09-21": "国民の休日",
  "2049-09-22": "秋分の日",
  "2049-10-11": "体育の日",
  "2049-11-03": "文化の日",
  "2049-11-23": "勤労感謝の日",
  "2050-01-01": "元日",
  "2050-01-10": "成人の日",
  "2050-02-11": "建国記念の日",
  "2050-02-23": "天皇誕生日",
  "2050-03-20": "春分の日",
  "2050-03-21": "振替休日",
  "2050-04-29": "昭和の日",
  "2050-05-03": "憲法記念日",
  "2050-05-04": "みどりの日",
  "2050-05-05": "こどもの日",
  "2050-07-18": "海の日",
  "2050-08-11": "山の日",
  "2050-09-19": "敬老の日",
  "2050-09-23": "秋分の日",
  "2050-10-10": "体育の日",
  "2050-11-03": "文化の日",
  "2050-11-23": "勤労感謝の日"
}
`
    return text
}

func holidaysMap() (map[string]string, error) {
    // Unmarshal json
    holidays := make(map[string]string)
    str := holidayJson()
    data := []byte(str)

    var f interface{}
    if err := json.Unmarshal(data, &f); err != nil {
        fmt.Println("JSON Unmarshal Error:", err)
        return holidays, err
    }

    // Covert to map[string]string
    m := f.(map[string]interface{})
    for k, v := range m {
        holidays[k] = fmt.Sprint(v)
    }

    return holidays, nil
}

func Prepare() {
    _holidays, _ = holidaysMap()
}

func IsHoliday(t time.Time) bool {
    var JST = time.FixedZone("JST", 3600*9)
    jst := t.In(JST)
    str := fmt.Sprintf("%04d-%02d-%02d", jst.Year(), jst.Month(), jst.Day())
    name := _holidays[str]
    return len(name) > 0
}
