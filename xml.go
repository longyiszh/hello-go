package main

import "fmt"
import "encoding/xml"

// the following xml is provided by @valorad/AOE3NativeED
var protos = []byte(`

<Units>

	<Unit id ='2100' name ='WCEJason'>
		<DBID>2100</DBID>
		<DisplayNameID>210000</DisplayNameID>
		<EditorNameID>230000</EditorNameID>
		<PopulationCount>1</PopulationCount>
		<ObstructionRadiusX>0.4900</ObstructionRadiusX>
		<ObstructionRadiusZ>0.4900</ObstructionRadiusZ>
		<FormationCategory>Body</FormationCategory>
		<MaxVelocity>6.0000</MaxVelocity>
		<MaxRunVelocity>8.0000</MaxRunVelocity>
		<MovementType>land</MovementType>
		<TurnRate>18.0000</TurnRate>
		<AnimFile>units\WCEJason\WCEJason.xml</AnimFile>
		<ImpactType>Flesh</ImpactType>
		<PhysicsInfo>dude</PhysicsInfo>
		<Icon>units\infantry\redolero\redolero_icon_64x64</Icon>
		<PortraitIcon>units\infantry\redolero\redolero_portrait</PortraitIcon>
		<RolloverTextID>250000</RolloverTextID>
		<ShortRolloverTextID>270000</ShortRolloverTextID>
		<InitialHitpoints>13500.0000</InitialHitpoints>
		<MaxHitpoints>13500.0000</MaxHitpoints>
		<LOS>12.0000</LOS>
		<UnitAIType>HandCombative</UnitAIType>
		<TrainPoints>300.0000</TrainPoints>
		<Bounty>1000.0000</Bounty>
		<BuildBounty>1000.0000</BuildBounty>
		<Cost resourcetype ='Food'>650.0000</Cost>
		<Cost resourcetype ='Gold'>350.0000</Cost>
		<AllowedAge>2</AllowedAge>
		<BuildLimit>1</BuildLimit>
		<Armor type ='Hand' value ='0.8000'></Armor>
		<UnitType>LogicalTypeHealed</UnitType>
		<UnitType>LogicalTypeValidSharpshoot</UnitType>
		<UnitType>LogicalTypeNeededForVictory</UnitType>
		<UnitType>LogicalTypeHandUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeLandMilitary</UnitType>
		<UnitType>LogicalTypeScout</UnitType>
		<UnitType>LogicalTypeGarrisonInShips</UnitType>
		<UnitType>LogicalTypeRangedUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeVillagersAttack</UnitType>
		<UnitType>LogicalTypeHandUnitsAttack</UnitType>
		<UnitType>LogicalTypeRangedUnitsAttack</UnitType>
		<UnitType>LogicalTypeMinimapFilterMilitary</UnitType>
		<UnitType>HasBountyValue</UnitType>
		<UnitType>CountsTowardMilitaryScore</UnitType>
		<UnitType>AbstractCavalryInfantry</UnitType>
		<UnitType>AbstractHandInfantry</UnitType>
		<UnitType>ConvertsHerds</UnitType>
		<UnitType>AbstractConsulateUnit</UnitType>
		<UnitType>UnitClass</UnitType>
		<UnitType>Military</UnitType>
		<UnitType>Unit</UnitType>
		<UnitType>Hero</UnitType>
		<UnitType>AbstractHeavyInfantry</UnitType>
		<UnitType>AbstractInfantry</UnitType>
		<Flag>NotDeleteable</Flag>
		<Flag>CollidesWithProjectiles</Flag>
		<Flag>ApplyHandicapTraining</Flag>
		<Flag>CorpseDecays</Flag>
		<Flag>ShowGarrisonButton</Flag>
		<Flag>DontRotateObstruction</Flag>
		<Flag>ObscuredByUnits</Flag>
		<Flag>Tracked</Flag>
		<Command page ='10' column ='1'>Stop</Command>
		<Command page ='10' column ='0'>Garrison</Command>
		<Command page ='11' column ='0'>Abilities</Command>
		<Tactics>WCEJason.tactics</Tactics>
		<ProtoAction>
			<Name>BuildingAttack</Name>
			<Damage>20.000000</Damage>
			<DamageType>Siege</DamageType>
			<ROF>3.000000</ROF>
		</ProtoAction>
		<ProtoAction>
			<Name>CoverBuildingAttack</Name>
			<Damage>10.000000</Damage>
			<DamageType>Siege</DamageType>
			<ROF>3.000000</ROF>
		</ProtoAction>
		<ProtoAction>
			<Name>CoverHandAttack</Name>
			<Damage>12.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.500000</ROF>
			<DamageBonus type ='AbstractCavalry'>7.00000</DamageBonus>
			<DamageBonus type ='AbstractLightInfantry'>4.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>DefendHandAttack</Name>
			<DamageBonus type ='AbstractLightInfantry'>4.50000</DamageBonus>
			<Damage>22.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.500000</ROF>
			<DamageBonus type ='AbstractCavalry'>7.00000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>MeleeHandAttack</Name>
			<Damage>22.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.500000</ROF>
			<DamageBonus type ='AbstractCavalry'>7.00000</DamageBonus>
			<DamageBonus type ='AbstractLightInfantry'>4.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>WCSShieldWallACT</Name>
			<MaxRange>1.000000</MaxRange>
		</ProtoAction>
	</Unit>

	<Unit id ='3100' name ='WCEShiriki'>
		<DBID>3100</DBID>
		<DisplayNameID>310000</DisplayNameID>
		<EditorNameID>330000</EditorNameID>
		<PopulationCount>1</PopulationCount>
		<ObstructionRadiusX>0.4900</ObstructionRadiusX>
		<ObstructionRadiusZ>0.4900</ObstructionRadiusZ>
		<FormationCategory>Body</FormationCategory>
		<MaxVelocity>6.0000</MaxVelocity>
		<MaxRunVelocity>8.0000</MaxRunVelocity>
		<MovementType>land</MovementType>
		<TurnRate>18.0000</TurnRate>
		<AnimFile>units\WCEShiriki\WCEShiriki.xml</AnimFile>
		<ImpactType>Flesh</ImpactType>
		<PhysicsInfo>dude</PhysicsInfo>
		<Icon>ui\units\az_skull_knight_icon</Icon>
		<PortraitIcon>ui\units\az_skull_knight_portrait</PortraitIcon>
		<RolloverTextID>350000</RolloverTextID>
		<ShortRolloverTextID>370000</ShortRolloverTextID>
		<InitialHitpoints>10000.0000</InitialHitpoints>
		<MaxHitpoints>10000.0000</MaxHitpoints>
		<LOS>20.0000</LOS>
		<UnitAIType>HandCombative</UnitAIType>
		<TrainPoints>300.0000</TrainPoints>
		<Bounty>1000.0000</Bounty>
		<BuildBounty>1000.0000</BuildBounty>
		<Cost resourcetype ='Gold'>1500.0000</Cost>
		<AllowedAge>2</AllowedAge>
		<BuildLimit>1</BuildLimit>
		<Armor type ='Hand' value ='0.2000'></Armor>
		<UnitType>LogicalTypeHealed</UnitType>
		<UnitType>LogicalTypeValidSharpshoot</UnitType>
		<UnitType>LogicalTypeNeededForVictory</UnitType>
		<UnitType>LogicalTypeHandUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeLandMilitary</UnitType>
		<UnitType>LogicalTypeScout</UnitType>
		<UnitType>LogicalTypeValidSPCUnitsDeadCondition</UnitType>
		<UnitType>LogicalTypeGarrisonInShips</UnitType>
		<UnitType>LogicalTypeRangedUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeVillagersAttack</UnitType>
		<UnitType>LogicalTypeHandUnitsAttack</UnitType>
		<UnitType>LogicalTypeRangedUnitsAttack</UnitType>
		<UnitType>LogicalTypeMinimapFilterMilitary</UnitType>
		<UnitType>HasBountyValue</UnitType>
		<UnitType>CountsTowardMilitaryScore</UnitType>
		<UnitType>AbstractCavalryInfantry</UnitType>
		<UnitType>ConvertsHerds</UnitType>
		<UnitType>AbstractHandInfantry</UnitType>
		<UnitType>MercType1</UnitType>
		<UnitType>Unit</UnitType>
		<UnitType>Hero</UnitType>
		<UnitType>UnitClass</UnitType>
		<UnitType>Military</UnitType>
		<UnitType>AbstractInfantry</UnitType>
		<UnitType>AbstractHeavyInfantry</UnitType>
		<Flag>CollidesWithProjectiles</Flag>
		<Flag>ApplyHandicapTraining</Flag>
		<Flag>CorpseDecays</Flag>
		<Flag>ShowGarrisonButton</Flag>
		<Flag>DontRotateObstruction</Flag>
		<Flag>ObscuredByUnits</Flag>
		<Flag>VisibleUnderFogIfGaia</Flag>
		<Flag>Tracked</Flag>
		<Flag>KnockoutDeath</Flag>
		<Flag>NotDeleteable</Flag>
		<Command page ='10' column ='1'>Stop</Command>
		<Command page ='10' column ='0'>Garrison</Command>
		<Command page ='11' column ='0'>Abilities</Command>
		<Tactics>WCEShiriki.tactics</Tactics>

		<ProtoAction>
			<Name>BuildingAttack</Name>
			<Damage>110.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>2.000000</ROF>
			<DamageCap>96.000000</DamageCap>
		</ProtoAction>
		<ProtoAction>
			<Name>CoverBuildingAttack</Name>
			<Damage>54.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>2.000000</ROF>
			<DamageCap>48.000000</DamageCap>
		</ProtoAction>

		<ProtoAction>
			<Name>CoverHandAttack</Name>
			<Damage>15.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.000000</ROF>
			<DamageCap>120.000000</DamageCap>
			<DamageBonus type ='AbstractCavalry'>6.000000</DamageBonus>
			<DamageArea>2.000000</DamageArea>
			<DamageFlags>Enemy</DamageFlags>
			<DamageBonus type ='AbstractLightInfantry'>4.000000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>DefendHandAttack</Name>
			<Damage>20.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.000000</ROF>
			<DamageCap>120.000000</DamageCap>
			<DamageBonus type ='AbstractCavalry'>6.000000</DamageBonus>
			<DamageArea>2.000000</DamageArea>
			<DamageFlags>Enemy</DamageFlags>
			<DamageBonus type ='AbstractLightInfantry'>4.000000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>GuardianAttack</Name>
			<Damage>30.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.000000</ROF>
		</ProtoAction>
		<ProtoAction>
			<Name>MeleeHandAttack</Name>
			<Damage>30.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>1.000000</ROF>
			<DamageCap>68.000000</DamageCap>
			<DamageBonus type ='AbstractCavalry'>6.000000</DamageBonus>
			<DamageArea>3.000000</DamageArea>
			<DamageFlags>Enemy</DamageFlags>
			<DamageBonus type ='AbstractLightInfantry'>4.000000</DamageBonus>
		</ProtoAction>
	</Unit>

	<Unit id ='4100' name ='WCEGongyi'>
		<DBID>4100</DBID>
		<DisplayNameID>410000</DisplayNameID>
		<EditorNameID>430000</EditorNameID>
		<PopulationCount>1</PopulationCount>
		<ObstructionRadiusX>0.4900</ObstructionRadiusX>
		<ObstructionRadiusZ>0.4900</ObstructionRadiusZ>
		<FormationCategory>Ranged</FormationCategory>
		<MaxVelocity>4.5000</MaxVelocity>
		<MaxRunVelocity>6.0000</MaxRunVelocity>
		<MovementType>land</MovementType>
		<TurnRate>18.0000</TurnRate>
		<AnimFile>units\WCEGongyi\WCEGongyi.xml</AnimFile>
		<ImpactType>Flesh</ImpactType>
		<PhysicsInfo>dude</PhysicsInfo>
		<Icon>units\asians\chinese\chu_ko_nu\chu_ko_nu_icon_64</Icon>
		<PortraitIcon>units\asians\chinese\chu_ko_nu\chu_ko_nu_icon_portrait</PortraitIcon>
		<RolloverTextID>450000</RolloverTextID>
		<ShortRolloverTextID>470000</ShortRolloverTextID>
		<InitialHitpoints>10000.0000</InitialHitpoints>
		<MaxHitpoints>10000.0000</MaxHitpoints>
		<LOS>20.0000</LOS>
		<ProjectileProtoUnit>Arrow</ProjectileProtoUnit>
		<AutoAttackRange>16.0000</AutoAttackRange>
		<UnitAIType>RangedCombative</UnitAIType>
		<TrainPoints>24.0000</TrainPoints>
		<Bounty>9.0000</Bounty>
		<BuildBounty>9.0000</BuildBounty>
		<Cost resourcetype ='Food'>355.0000</Cost>
		<Cost resourcetype ='Gold'>555.0000</Cost>
		<AllowedAge>2</AllowedAge>
		<BuildLimit>1</BuildLimit>
		<Armor type ='Ranged' value ='0.2000'></Armor>
		<UnitType>LogicalTypeHealed</UnitType>
		<UnitType>LogicalTypeValidSharpshoot</UnitType>
		<UnitType>LogicalTypeNeededForVictory</UnitType>
		<UnitType>LogicalTypeHandUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeLandMilitary</UnitType>
		<UnitType>LogicalTypeScout</UnitType>
		<UnitType>LogicalTypeValidSPCUnitsDeadCondition</UnitType>
		<UnitType>LogicalTypeGarrisonInShips</UnitType>
		<UnitType>LogicalTypeRangedUnitsAutoAttack</UnitType>
		<UnitType>LogicalTypeVillagersAttack</UnitType>
		<UnitType>LogicalTypeHandUnitsAttack</UnitType>
		<UnitType>LogicalTypeRangedUnitsAttack</UnitType>
		<UnitType>LogicalTypeMinimapFilterMilitary</UnitType>
		<UnitType>ConvertsHerds</UnitType>
		<UnitType>AbstractRangedInfantry</UnitType>
		<UnitType>Ranged</UnitType>
		<UnitType>HasBountyValue</UnitType>
		<UnitType>AbstractCavalryInfantry</UnitType>
		<UnitType>CountsTowardMilitaryScore</UnitType>
		<UnitType>UnitClass</UnitType>
		<UnitType>Military</UnitType>
		<UnitType>Unit</UnitType>
		<UnitType>Hero</UnitType>
		<UnitType>AbstractArcher</UnitType>
		<UnitType>AbstractInfantry</UnitType>
		<Flag>CollidesWithProjectiles</Flag>
		<Flag>ApplyHandicapTraining</Flag>
		<Flag>CorpseDecays</Flag>
		<Flag>ShowGarrisonButton</Flag>
		<Flag>DontRotateObstruction</Flag>
		<Flag>ObscuredByUnits</Flag>
		<Flag>Tracked</Flag>
		<Flag>NotDeleteable</Flag>
		<Flag>KnockoutDeath</Flag>
		<Command page ='10' column ='1'>Stop</Command>
		<Command page ='10' column ='0'>Garrison</Command>
		<Command page ='11' column ='1'>Abilities</Command>
		<Tactics>WCEGongyi.tactics</Tactics>
		<ProtoAction>
			<Name>BuildingAttack</Name>
			<Damage>6.000000</Damage>
			<DamageType>Siege</DamageType>
			<ROF>1.000000</ROF>
		</ProtoAction>
		<ProtoAction>
			<Name>DefendHandAttack</Name>
			<Damage>12.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>0.500000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>DefendRangedAttack</Name>
			<Damage>10.000000</Damage>
			<DamageType>Ranged</DamageType>
			<MinRange>2.000000</MinRange>
			<MaxRange>16.000000</MaxRange>
			<ROF>1.000000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>MeleeHandAttack</Name>
			<Damage>12.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>0.500000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>StaggerHandAttack</Name>
			<Damage>12.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>0.500000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>StaggerRangedAttack</Name>
			<Damage>10.000000</Damage>
			<DamageType>Ranged</DamageType>
			<MinRange>2.000000</MinRange>
			<MaxRange>16.000000</MaxRange>
			<ROF>1.000000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>VolleyHandAttack</Name>
			<Damage>12.000000</Damage>
			<DamageType>Hand</DamageType>
			<ROF>0.500000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
		</ProtoAction>
		<ProtoAction>
			<Name>VolleyRangedAttack</Name>
			<Damage>10.000000</Damage>
			<DamageType>Ranged</DamageType>
			<MinRange>2.000000</MinRange>
			<MaxRange>16.000000</MaxRange>
			<ROF>1.000000</ROF>
			<DamageBonus type ='AbstractLightCavalry'>4.000000</DamageBonus>
			<DamageBonus type ='xpEagleKnight'>4.000000</DamageBonus>
			<DamageBonus type ='AbstractHeavyInfantry'>4.000000</DamageBonus>
			<DamageBonus type ='xpCoyoteMan'>1.50000</DamageBonus>
			<DamageBonus type ='AbstractCavalry'>1.50000</DamageBonus>
		</ProtoAction>
	</Unit>

</Units>
`)

// Protos {<Units>...</Units>}
type Protos struct {
	Units []Unit `xml:"Unit"`
}

// Unit {<Unit>...</Unit>}
type Unit struct {
	DBID       int    `xml:"DBID"`
	AllowedAge int    `xml:"AllowedAge"`
	Tactics    string `xml:"Tactics"`
}

// UnitAction {<ProtoAction>...</ProtoAction><ProtoAction>...</ProtoAction><ProtoAction>...</ProtoAction>}
// xml nested-tag location navigation
type UnitAction struct {
	Actions []Action `xml:"Unit>ProtoAction"`
}

// Action {<ProtoAction>...</ProtoAction>}
type Action struct {
	Name       string  `xml:"Name"`
	Damage     float32 `xml:"Damage"`
	DamageType string  `xml:"DamageType"`
}

// func (u Unit) String() string {
// 	return fmt.Sprintf(u.Tactics)
// }

func decodeProtos() Protos {
	var pr Protos
	xml.Unmarshal(protos, &pr)

	return pr
}

func decodeAction() UnitAction {
	var act UnitAction
	xml.Unmarshal(protos, &act)

	return act
}

func mainXML() {
	// fmt.Println("xml")
	fmt.Println(decodeAction())
	// units := decodeProtos().Units

	// for index, unit := range units {
	// 	fmt.Printf("\n\n%d-%+v", index, unit)
	// 	fmt.Printf("\n%d-%s", index, unit.Tactics)
	// }

}

// note:
// [5]type <-- array: fixed length 5
// []type <-- slice
