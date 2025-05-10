package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int `json:"code"`
	Message any `json:"msg"`
}

func main() {
	config := GetConfig()
	app := echo.New()

	a := `<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:u="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	    <s:Header>
	        <a:Action s:mustUnderstand="1">http://www.microsoft.com/SoftwareDistribution/Server/ClientWebService/SyncUpdatesResponse</a:Action>
	        <a:RelatesTo>urn:uuid:175df68c-4b91-41ee-b70b-f2208c65438e</a:RelatesTo>
	        <o:Security s:mustUnderstand="1" xmlns:o="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
	            <u:Timestamp u:Id="_0">
	                <u:Created>2025-05-01T06:49:22.002Z</u:Created>
	                <u:Expires>2025-05-01T06:54:22.002Z</u:Expires>
	            </u:Timestamp>
	        </o:Security>
	    </s:Header>
	    <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	        <SyncUpdatesResponse xmlns="http://www.microsoft.com/SoftwareDistribution/Server/ClientWebService">
	            <SyncUpdatesResult>
	                <NewUpdates>
	                    <UpdateInfo>
	                        <ID>316003061</ID>
	                        <Deployment>
	                            <ID>535903338</ID>
	                            <Action>Evaluate</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-23</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>false</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="dd78b8a1-0b20-45c1-add6-4da72e9364cf" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Category" /&gt;&lt;Relationships /&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;b.WindowsVersion Comparison="GreaterThanOrEqualTo" MajorVersion="10" MinorVersion="0" /&gt;&lt;/IsInstalled&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>296374060</ID>
	                        <Deployment>
	                            <ID>471931323</ID>
	                            <Action>Evaluate</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-03-14</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>false</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="E0789628-CE08-4437-BE74-2495B842F43B" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Category" /&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;True /&gt;&lt;/IsInstalled&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>315728269</ID>
	                        <Deployment>
	                            <ID>534654284</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-16</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="2aeac134-a06a-428d-a35f-c42e591b73cd" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" /&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" /&gt;&lt;/AtLeastOne&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="cd5ffd1e-e932-4e3a-bf74-18bf0b1bbd83" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;BundledUpdates&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="2f384bb5-35f3-466f-ab24-e383959ef5a6" RevisionNumber="1" /&gt;&lt;UpdateIdentity UpdateID="ae6fcf00-bc10-4f1a-b9af-ece5238870e1" RevisionNumber="1" /&gt;&lt;UpdateIdentity UpdateID="b984e085-e1b6-4164-a450-24df084a9396" RevisionNumber="1" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c9007f36-a00e-4cf3-8a08-b0fe1bef1ecf" RevisionNumber="1" /&gt;&lt;/BundledUpdates&gt;&lt;/Relationships&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>315728270</ID>
	                        <Deployment>
	                            <ID>534654285</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-16</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="2f384bb5-35f3-466f-ab24-e383959ef5a6" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" PackageRank="30001"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="e7cd787c-8594-4e6b-8302-174f4aed7a72" /&gt;&lt;UpdateIdentity UpdateID="8a384ffa-f490-45be-a936-4b9891e74349" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.MinecraftUWP_1.21.7301.0_x86__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isMain":false,"content.packageId":"Microsoft.MinecraftUWP_1.21.7301.0_x86__8wekyb3d8bbwe","content.productId":"e90df8bd-dc71-43be-b7ef-648205f09325","content.targetPlatforms":[{"platform.maxVersionTested":2814751014977536,"platform.minVersion":2814751014977536,"platform.target":0}],"content.type":7,"policy":{"category.first":"game","category.second":"Games","category.third":"Action &amp;amp; adventure","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":2,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":17,"systemId":6},{"level":55,"systemId":13},{"level":8,"systemId":3},{"level":13,"systemId":5},{"level":48,"systemId":12},{"level":28,"systemId":9},{"level":77,"systemId":16},{"level":40,"systemId":10},{"level":70,"systemId":15}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>315728271</ID>
	                        <Deployment>
	                            <ID>534654286</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-16</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="ae6fcf00-bc10-4f1a-b9af-ece5238870e1" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" PackageRank="30002"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="8a384ffa-f490-45be-a936-4b9891e74349" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.MinecraftUWP_1.21.7301.0_x64__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isMain":false,"content.packageId":"Microsoft.MinecraftUWP_1.21.7301.0_x64__8wekyb3d8bbwe","content.productId":"e90df8bd-dc71-43be-b7ef-648205f09325","content.targetPlatforms":[{"platform.maxVersionTested":2814751014977536,"platform.minVersion":2814751014977536,"platform.target":0}],"content.type":7,"policy":{"category.first":"game","category.second":"Games","category.third":"Action &amp;amp; adventure","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":2,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":17,"systemId":6},{"level":55,"systemId":13},{"level":8,"systemId":3},{"level":13,"systemId":5},{"level":48,"systemId":12},{"level":28,"systemId":9},{"level":77,"systemId":16},{"level":40,"systemId":10},{"level":70,"systemId":15}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520318</ID>
	                        <Deployment>
	                            <ID>475529676</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="c9007f36-a00e-4cf3-8a08-b0fe1bef1ecf" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" IsAppxFramework="true" /&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="cd5ffd1e-e932-4e3a-bf74-18bf0b1bbd83" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;BundledUpdates&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="6194cec0-df15-4b08-af05-b09565805255" RevisionNumber="1" /&gt;&lt;UpdateIdentity UpdateID="addca3d4-7d1c-4b6e-b9d4-95ba502f65de" RevisionNumber="1" /&gt;&lt;UpdateIdentity UpdateID="c3478097-f09c-4622-a8bf-6b1a5ba7db12" RevisionNumber="1" /&gt;&lt;UpdateIdentity UpdateID="ca9a3dd6-defd-48d4-abfc-f5c5c2b7a10a" RevisionNumber="1" /&gt;&lt;/AtLeastOne&gt;&lt;/BundledUpdates&gt;&lt;/Relationships&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520322</ID>
	                        <Deployment>
	                            <ID>475529680</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="ca9a3dd6-defd-48d4-abfc-f5c5c2b7a10a" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" IsAppxFramework="true" PackageRank="100"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.VCLibs.140.00_14.0.33519.0_arm64__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isFramework":true,"content.isMain":false,"content.packageId":"Microsoft.VCLibs.140.00_14.0.33519.0_arm64__8wekyb3d8bbwe","content.productId":"938c94d3-76d8-49d2-8524-dd6e4581ff3a","content.targetPlatforms":[{"platform.maxVersionTested":2814750425219072,"platform.minVersion":2814750425219072,"platform.target":0}],"content.type":7,"policy":{"category.first":"app","category.second":"Developer tools","category.third":"Development kits","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":1,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":7,"systemId":3},{"level":12,"systemId":5},{"level":48,"systemId":12},{"level":27,"systemId":9},{"level":76,"systemId":16},{"level":68,"systemId":15},{"level":54,"systemId":13}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520319</ID>
	                        <Deployment>
	                            <ID>475529677</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="6194cec0-df15-4b08-af05-b09565805255" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" IsAppxFramework="true" PackageRank="100"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="8a384ffa-f490-45be-a936-4b9891e74349" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.VCLibs.140.00_14.0.33519.0_x64__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isFramework":true,"content.isMain":false,"content.packageId":"Microsoft.VCLibs.140.00_14.0.33519.0_x64__8wekyb3d8bbwe","content.productId":"938c94d3-76d8-49d2-8524-dd6e4581ff3a","content.targetPlatforms":[{"platform.maxVersionTested":2814750425219072,"platform.minVersion":2814750425219072,"platform.target":0}],"content.type":7,"policy":{"category.first":"app","category.second":"Developer tools","category.third":"Development kits","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":1,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":7,"systemId":3},{"level":12,"systemId":5},{"level":48,"systemId":12},{"level":27,"systemId":9},{"level":76,"systemId":16},{"level":68,"systemId":15},{"level":54,"systemId":13}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520321</ID>
	                        <Deployment>
	                            <ID>475529679</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="c3478097-f09c-4622-a8bf-6b1a5ba7db12" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" IsAppxFramework="true" PackageRank="100"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="3d854509-0f65-46d6-b804-83b3e28c47e0" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.VCLibs.140.00_14.0.33519.0_arm__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isFramework":true,"content.isMain":false,"content.packageId":"Microsoft.VCLibs.140.00_14.0.33519.0_arm__8wekyb3d8bbwe","content.productId":"938c94d3-76d8-49d2-8524-dd6e4581ff3a","content.targetPlatforms":[{"platform.maxVersionTested":2814750425219072,"platform.minVersion":2814750425219072,"platform.target":0}],"content.type":7,"policy":{"category.first":"app","category.second":"Developer tools","category.third":"Development kits","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":1,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":7,"systemId":3},{"level":12,"systemId":5},{"level":48,"systemId":12},{"level":27,"systemId":9},{"level":76,"systemId":16},{"level":68,"systemId":15},{"level":54,"systemId":13}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520320</ID>
	                        <Deployment>
	                            <ID>475529678</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="addca3d4-7d1c-4b6e-b9d4-95ba502f65de" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" IsAppxFramework="true" PackageRank="100"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="e7cd787c-8594-4e6b-8302-174f4aed7a72" /&gt;&lt;UpdateIdentity UpdateID="8a384ffa-f490-45be-a936-4b9891e74349" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.VCLibs.140.00_14.0.33519.0_x86__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isFramework":true,"content.isMain":false,"content.packageId":"Microsoft.VCLibs.140.00_14.0.33519.0_x86__8wekyb3d8bbwe","content.productId":"938c94d3-76d8-49d2-8524-dd6e4581ff3a","content.targetPlatforms":[{"platform.maxVersionTested":2814750425219072,"platform.minVersion":2814750425219072,"platform.target":0}],"content.type":7,"policy":{"category.first":"app","category.second":"Developer tools","category.third":"Development kits","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":1,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":7,"systemId":3},{"level":12,"systemId":5},{"level":48,"systemId":12},{"level":27,"systemId":9},{"level":76,"systemId":16},{"level":68,"systemId":15},{"level":54,"systemId":13}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>315728268</ID>
	                        <Deployment>
	                            <ID>534654288</ID>
	                            <Action>Install</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-16</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="b5be71c9-79cd-467c-91c6-eded1404f4ee" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" ExplicitlyDeployable="true" PerUser="true" ApplyPackageRank="true" /&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" /&gt;&lt;/AtLeastOne&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="cd5ffd1e-e932-4e3a-bf74-18bf0b1bbd83" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;BundledUpdates&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="2aeac134-a06a-428d-a35f-c42e591b73cd" RevisionNumber="1" /&gt;&lt;/AtLeastOne&gt;&lt;/BundledUpdates&gt;&lt;/Relationships&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>297520317</ID>
	                        <Deployment>
	                            <ID>475529681</ID>
	                            <Action>Install</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2024-04-10</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="7c4b0956-b3f1-40ac-a633-668abc5d6af2" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" ExplicitlyDeployable="true" PerUser="true" IsAppxFramework="true" ApplyPackageRank="true" /&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="c746f7a0-3385-4c03-b4fa-7da2943db294" /&gt;&lt;/AtLeastOne&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="cd5ffd1e-e932-4e3a-bf74-18bf0b1bbd83" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="90cfdd83-c886-4525-bd71-4981e4adb45c" /&gt;&lt;/Prerequisites&gt;&lt;BundledUpdates&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="c9007f36-a00e-4cf3-8a08-b0fe1bef1ecf" RevisionNumber="1" /&gt;&lt;/AtLeastOne&gt;&lt;/BundledUpdates&gt;&lt;/Relationships&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>315728272</ID>
	                        <Deployment>
	                            <ID>534654287</ID>
	                            <Action>Bundle</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2025-04-16</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>true</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="b984e085-e1b6-4164-a450-24df084a9396" RevisionNumber="1" /&gt;&lt;Properties UpdateType="Software" PerUser="true" PackageRank="30003"&gt;&lt;SecuredFragment&gt;FileUrl&lt;/SecuredFragment&gt;&lt;/Properties&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" /&gt;&lt;/AtLeastOne&gt;&lt;UpdateIdentity UpdateID="31ec4b55-3ef0-4a75-aca2-18c7180637b7" /&gt;&lt;AtLeastOne&gt;&lt;UpdateIdentity UpdateID="3d854509-0f65-46d6-b804-83b3e28c47e0" /&gt;&lt;UpdateIdentity UpdateID="25ce6bbd-894d-486a-b7d2-1bd23f192cfe" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;IsInstallable&gt;&lt;AppxPackageInstallable /&gt;&lt;/IsInstallable&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxMetadata PackageType="UAP" IsAppxBundle="false" PackageMoniker="Microsoft.MinecraftUWP_1.21.7301.0_arm__8wekyb3d8bbwe"&gt;&lt;ApplicabilityBlob&gt;{"blob.version":1688867040526336,"content.isMain":false,"content.packageId":"Microsoft.MinecraftUWP_1.21.7301.0_arm__8wekyb3d8bbwe","content.productId":"e90df8bd-dc71-43be-b7ef-648205f09325","content.targetPlatforms":[{"platform.maxVersionTested":2814751014977536,"platform.minVersion":2814751014977536,"platform.target":0}],"content.type":7,"policy":{"category.first":"game","category.second":"Games","category.third":"Action &amp;amp; adventure","optOut.backupRestore":false,"optOut.removeableMedia":false},"policy2":{"ageRating":2,"optOut.DVR":false,"thirdPartyAppRatings":[{"level":17,"systemId":6},{"level":55,"systemId":13},{"level":8,"systemId":3},{"level":13,"systemId":5},{"level":48,"systemId":12},{"level":28,"systemId":9},{"level":77,"systemId":16},{"level":40,"systemId":10},{"level":70,"systemId":15}]}}&lt;/ApplicabilityBlob&gt;&lt;/AppxMetadata&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                    <UpdateInfo>
	                        <ID>298196079</ID>
	                        <Deployment>
	                            <ID>135495451</ID>
	                            <Action>Evaluate</Action>
	                            <IsAssigned>true</IsAssigned>
	                            <LastChangeTime>2016-03-25</LastChangeTime>
	                        </Deployment>
	                        <IsLeaf>false</IsLeaf>
	                        <Xml>&lt;UpdateIdentity UpdateID="d25480ca-36aa-46e6-b76b-39608d49558c" RevisionNumber="2" /&gt;&lt;Properties UpdateType="Category" PerUser="true" /&gt;&lt;Relationships&gt;&lt;Prerequisites&gt;&lt;AtLeastOne IsCategory="true"&gt;&lt;UpdateIdentity UpdateID="1e16fab0-b667-4a09-ad2a-497f61d6280c" /&gt;&lt;/AtLeastOne&gt;&lt;/Prerequisites&gt;&lt;/Relationships&gt;&lt;ApplicabilityRules&gt;&lt;IsInstalled&gt;&lt;AppxPackageInstalled /&gt;&lt;/IsInstalled&gt;&lt;Metadata&gt;&lt;AppxPackageMetadata&gt;&lt;AppxFamilyMetadata Name="MICROSOFT.MINECRAFTUWP" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" LegacyMobileProductId="e90df8bd-dc71-43be-b7ef-648205f09325" /&gt;&lt;/AppxPackageMetadata&gt;&lt;/Metadata&gt;&lt;/ApplicabilityRules&gt;</Xml>
	                    </UpdateInfo>
	                </NewUpdates>
	                <Truncated>false</Truncated>
	                <NewCookie>
	                    <Expiration>0028-08-07T00:00:00Z</Expiration>
	                    <EncryptedData>ATVNqfLH9EAdrZGpzH5KSedbV0mM00ZPuYIvvb8d+YLOmgmn820Xow7saqfOsKXiDkAXhNkHxYcdrAt56LYL6HrSYcZtRS1xBsfuzDja8rhJQp8N9QWv0wHlCzMysG7z18R0qzZfXiP5ek7Q2lsSclnxAyHzER9P7J2g/lbjSi7AWP0KGO77T8KcfLqN3uvunpb/mZ6auOKFTSsKn53oNhI9iowxz274EmPjtNjiyiQxshO23Yn1wS96v9swG4GCag==</EncryptedData>
	                </NewCookie>
	                <ExtendedUpdateInfo>
	                    <Updates>
	                        <Update>
	                            <ID>297520317</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" IsAppxFramework="true" FromStoreService="true" PackageIdentityName="Microsoft.VCLibs.140.00" LegacyMobileProductId="938c94d3-76d8-49d2-8524-dd6e4581ff3a" /&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>315728268</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" CreationDate="2025-04-16T00:37:14.0282216Z" ContentType="Application" IsAppxFramework="false" FromStoreService="true" PackageIdentityName="MICROSOFT.MINECRAFTUWP" LegacyMobileProductId="e90df8bd-dc71-43be-b7ef-648205f09325" /&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>297520320</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="true" DefaultPropertiesLanguage="en" MaxDownloadSize="760035" MinDownloadSize="0" FromStoreService="true" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" PackageIdentityName="Microsoft.VCLibs.140.00" PackageContentId="d5341d97-2f34-3aa0-c0d0-b693b8541db6"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="9683459a-02fa-4bd6-9ae6-af8ddfbeef35.appx" Digest="/DWIkZI6XJwxOY/s/GAOyxuZIBQ=" DigestAlgorithm="SHA1" Size="758544" Modified="2024-01-29T20:03:05.7437236Z" InstallerSpecificIdentifier="Microsoft.VCLibs.140.00_14.0.33519.0_x86__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;e6bqe8Ms1Yt+BoPaWIeWCGrM+3Tvt6PlJen4AU0q1mM=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;xF5p1A9LiL1aEiYWHG7IxoBzYLH6fTxXC8TTtH3/7NM=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;xPnH1+pblf7oK+ivWJwv1GUym6aXbAIuBqJq02SrG/c=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_9683459a-02fa-4bd6-9ae6-af8ddfbeef35.cab" Digest="R3wKJD2qNYxtduw50Ezcxv6Blg4=" DigestAlgorithm="SHA1" Size="1491" Modified="2024-01-29T20:03:05.7437236Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;xPnH1+pblf7oK+ivWJwv1GUym6aXbAIuBqJq02SrG/c=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="9683459a-02fa-4bd6-9ae6-af8ddfbeef35.appx" MainPackage="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>297520321</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="true" DefaultPropertiesLanguage="en" MaxDownloadSize="837201" MinDownloadSize="0" FromStoreService="true" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" PackageIdentityName="Microsoft.VCLibs.140.00" PackageContentId="d5341d97-2f34-3aa0-c0d0-b693b8541db6"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="f52e53fc-9664-4720-800e-a637f6f5c125.appx" Digest="QcvoJerwPZqVHhAet72YUZvX/eI=" DigestAlgorithm="SHA1" Size="835614" Modified="2024-01-29T20:03:05.7437236Z" InstallerSpecificIdentifier="Microsoft.VCLibs.140.00_14.0.33519.0_arm__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;LEIlI/9pNomoTBCVhc+kRBQ6w7a3pcrfSFivxqPLdQ8=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;0jxxtMWYZUGx9PjKxcKfR8dVl1ZKXdkA/qFW3WPNv3k=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;wFiAGZ3fuldahbUjns4+RksTqg9WgTi6v91vsDQBG+w=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_f52e53fc-9664-4720-800e-a637f6f5c125.cab" Digest="eZ9xNey/fo1yKnfdjbPUEVpn0JY=" DigestAlgorithm="SHA1" Size="1587" Modified="2024-01-29T20:03:05.7437236Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;wFiAGZ3fuldahbUjns4+RksTqg9WgTi6v91vsDQBG+w=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="f52e53fc-9664-4720-800e-a637f6f5c125.appx" MainPackage="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>297520319</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="true" DefaultPropertiesLanguage="en" MaxDownloadSize="898248" MinDownloadSize="0" FromStoreService="true" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" PackageIdentityName="Microsoft.VCLibs.140.00" PackageContentId="d5341d97-2f34-3aa0-c0d0-b693b8541db6"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="c452d4ef-2486-4efe-9c99-36b3d23e0160.appx" Digest="AMWhizJDyZKWck1MApdbqPw/81M=" DigestAlgorithm="SHA1" Size="896581" Modified="2024-01-29T20:03:05.7437236Z" InstallerSpecificIdentifier="Microsoft.VCLibs.140.00_14.0.33519.0_x64__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;nBe1IfnWkKH1BNpRCO1u7FZp6zqP0TMe70PkDYTnQoM=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;jJGN8nOlOnu8YQbY0R7R4HbWi84G2xoW6PpRAYJd/fw=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;EonSpJ2OVmkTezWws2DAwPoxif4BKC5VqMCSWGhbEtE=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_c452d4ef-2486-4efe-9c99-36b3d23e0160.cab" Digest="qfRBBNC3ihHP3svRdwmAB3PpIa4=" DigestAlgorithm="SHA1" Size="1667" Modified="2024-01-29T20:03:05.7437236Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;EonSpJ2OVmkTezWws2DAwPoxif4BKC5VqMCSWGhbEtE=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="c452d4ef-2486-4efe-9c99-36b3d23e0160.appx" MainPackage="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>297520322</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="true" DefaultPropertiesLanguage="en" MaxDownloadSize="1572933" MinDownloadSize="0" FromStoreService="true" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" PackageIdentityName="Microsoft.VCLibs.140.00" PackageContentId="d5341d97-2f34-3aa0-c0d0-b693b8541db6"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="4aa2f4c4-ecd2-41d1-8089-304a95524359.appx" Digest="FNsrOe2gPtP09mVAhY+up+uMz3Y=" DigestAlgorithm="SHA1" Size="1570406" Modified="2024-01-29T20:03:05.7437236Z" InstallerSpecificIdentifier="Microsoft.VCLibs.140.00_14.0.33519.0_arm64__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;xz0PVd2jMfncvvyZ/1pCC2ISB3PSkXOHY5OCqkeFM+4=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;hLClcwdkDyNPHdpu7XlaqBz77hT6bQiO3ZCrLE+8MLg=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;tADpKYksLidkMgiCFJbEtUU7JCcG6Qtr72t4YNunJOM=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_4aa2f4c4-ecd2-41d1-8089-304a95524359.cab" Digest="PL2nid7mClkBJLHWEB8UEZdm4k8=" DigestAlgorithm="SHA1" Size="2527" Modified="2024-01-29T20:03:05.7437236Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;tADpKYksLidkMgiCFJbEtUU7JCcG6Qtr72t4YNunJOM=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="4aa2f4c4-ecd2-41d1-8089-304a95524359.appx" MainPackage="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>297520318</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" CreationDate="2024-04-10T21:34:54.6743066Z" ContentType="Application" IsAppxFramework="true" FromStoreService="true" PackageIdentityName="Microsoft.VCLibs.140.00" LegacyMobileProductId="938c94d3-76d8-49d2-8524-dd6e4581ff3a" /&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>315728271</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="false" DefaultPropertiesLanguage="en" MaxDownloadSize="861397689" MinDownloadSize="0" FromStoreService="true" CreationDate="2025-04-16T00:37:14.0282216Z" ContentType="Application" PackageIdentityName="MICROSOFT.MINECRAFTUWP" PackageContentId="d5e7c859-f35a-bfee-9934-71b7a9760947"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="dc0e0caf-5a1f-4b6d-aede-d53387f0c9b7.appx" Digest="YP5WNN55k5GfWkFmcM0uJY5JeDk=" DigestAlgorithm="SHA1" Size="860376012" Modified="2025-04-11T22:54:50.4349755Z" InstallerSpecificIdentifier="Microsoft.MinecraftUWP_1.21.7301.0_x64__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;JCeZCdIkvVADv0BzurW4kKig7K/F2pPYlTCOEDRanxs=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;2+Za3uQnhxfieDOe8ulV8QnaLyAi9KmYaLmRzbC1xtc=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;Mal2H4ef3+Av1NVWEBchtD/vsT/Vuc7U4k2369reTyE=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_dc0e0caf-5a1f-4b6d-aede-d53387f0c9b7.cab" Digest="H4DZLjswsK1FAVjLo2bS8TxdTtE=" DigestAlgorithm="SHA1" Size="1021677" Modified="2025-04-11T22:54:50.4349755Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;Mal2H4ef3+Av1NVWEBchtD/vsT/Vuc7U4k2369reTyE=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="dc0e0caf-5a1f-4b6d-aede-d53387f0c9b7.appx" MainPackage="true" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>315728270</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="false" DefaultPropertiesLanguage="en" MaxDownloadSize="844654858" MinDownloadSize="0" FromStoreService="true" CreationDate="2025-04-16T00:37:14.0282216Z" ContentType="Application" PackageIdentityName="MICROSOFT.MINECRAFTUWP" PackageContentId="d5e7c859-f35a-bfee-9934-71b7a9760947"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="3d168df5-eaa5-4478-b5f7-7e7594a8d713.appx" Digest="2NcJwRI7skl6+7Grr06gZBuf0TQ=" DigestAlgorithm="SHA1" Size="843653271" Modified="2025-04-11T22:54:50.3724650Z" InstallerSpecificIdentifier="Microsoft.MinecraftUWP_1.21.7301.0_x86__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;oiSeFzQ12jTOSjg0hDBTlzL6UN/YZzc8kBbatvW0BnE=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;/8QdpFzdiNy6t3HJ/u/ePOuZBWO1C9/P9KpviDxemkk=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;rbF6In1VmciVzs7kQGOKNgmQnCiw8KW/6jOr2vidyq4=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_3d168df5-eaa5-4478-b5f7-7e7594a8d713.cab" Digest="15NnHGqv2+72D/xnPqCjN8a8b0I=" DigestAlgorithm="SHA1" Size="1001587" Modified="2025-04-11T22:54:50.3724650Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;rbF6In1VmciVzs7kQGOKNgmQnCiw8KW/6jOr2vidyq4=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="3d168df5-eaa5-4478-b5f7-7e7594a8d713.appx" MainPackage="true" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>315728269</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" CreationDate="2025-04-16T00:37:14.0282216Z" ContentType="Application" IsAppxFramework="false" FromStoreService="true" PackageIdentityName="MICROSOFT.MINECRAFTUWP" LegacyMobileProductId="e90df8bd-dc71-43be-b7ef-648205f09325" /&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>296374060</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/Category" CreationDate="2024-03-06T18:00:00.000Z" /&gt;&lt;HandlerSpecificData type="cat:Category"&gt;&lt;CategoryInformation CategoryType="UpdateClassification" ProhibitsSubcategories="true" ProhibitsUpdates="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>316003061</ID>
	                            <Xml>&lt;ExtendedProperties DefaultPropertiesLanguage="en" Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/Category" CreationDate="2015-08-25T20:09:50.002Z" /&gt;&lt;HandlerSpecificData type="cat:Category"&gt;&lt;CategoryInformation CategoryType="Product" ProhibitsSubcategories="true" ExcludedByDefault="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>298196079</ID>
	                            <Xml>&lt;ExtendedProperties IsAppxFramework="false" Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/Category" DefaultPropertiesLanguage="en" FromStoreService="true" ContentType="Application" /&gt;&lt;HandlerSpecificData type="cat:Category"&gt;&lt;CategoryInformation CategoryType="Application" ProhibitsSubcategories="true" ExcludedByDefault="false" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                        <Update>
	                            <ID>315728272</ID>
	                            <Xml>&lt;ExtendedProperties Handler="http://schemas.microsoft.com/msus/2002/12/UpdateHandlers/AppxPackage" IsAppxFramework="false" DefaultPropertiesLanguage="en" MaxDownloadSize="838138466" MinDownloadSize="0" FromStoreService="true" CreationDate="2025-04-16T00:37:14.0282216Z" ContentType="Application" PackageIdentityName="MICROSOFT.MINECRAFTUWP" PackageContentId="d5e7c859-f35a-bfee-9934-71b7a9760947"&gt;&lt;InstallationBehavior /&gt;&lt;/ExtendedProperties&gt;&lt;Files&gt;&lt;File FileName="832d0da6-9cdd-467d-842a-9bbe728ead1b.appx" Digest="iKx3kXVgixSuSl1HH9aIu/Nk8hg=" DigestAlgorithm="SHA1" Size="837145161" Modified="2025-04-11T22:54:50.4974666Z" InstallerSpecificIdentifier="Microsoft.MinecraftUWP_1.21.7301.0_arm__8wekyb3d8bbwe"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;+A7DT3Oeb7hW7FUrNq6QzZUFkUA0HA8g+2SNu1u0r9E=&lt;/AdditionalDigest&gt;&lt;PiecesHashDigest Algorithm="SHA256"&gt;gIKRIwsXUEJP4y99XHnoPw71N0loNpwOZPRIOjGUJEE=&lt;/PiecesHashDigest&gt;&lt;BlockMapDigest Algorithm="SHA256"&gt;/NqRo6R6Gv9ZsH73udVjwtRl0qb/N6JO5HmB4Gm7naE=&lt;/BlockMapDigest&gt;&lt;/File&gt;&lt;File FileName="Abm_832d0da6-9cdd-467d-842a-9bbe728ead1b.cab" Digest="nR1AQz0UUC0CXSNS3QlDR4eI7m8=" DigestAlgorithm="SHA1" Size="993305" Modified="2025-04-11T22:54:50.4974666Z" PatchingType="DynamicMetadata"&gt;&lt;AdditionalDigest Algorithm="SHA256"&gt;/NqRo6R6Gv9ZsH73udVjwtRl0qb/N6JO5HmB4Gm7naE=&lt;/AdditionalDigest&gt;&lt;/File&gt;&lt;/Files&gt;&lt;HandlerSpecificData type="appx:AppxInstaller"&gt;&lt;AppxPackageInstallData PackageFileName="832d0da6-9cdd-467d-842a-9bbe728ead1b.appx" MainPackage="true" /&gt;&lt;/HandlerSpecificData&gt;</Xml>
	                        </Update>
	                    </Updates>
	                </ExtendedUpdateInfo>
	            </SyncUpdatesResult>
	        </SyncUpdatesResponse>
	    </s:Body>
	</s:Envelope>`

	log.Print(html.UnescapeString(a))

	app.GET("/*", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "Microsoft Store API",
		})
	})

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", config.Port)))
}
