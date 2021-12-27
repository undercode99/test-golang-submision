SELECT 
    usr."ID", usr."UserName", prnt."UserName" as "ParentUserName"
FROM 
    "USER" as usr left join "USER" as prnt on prnt."ID"=usr."Parent" 
order by usr."ID" asc
