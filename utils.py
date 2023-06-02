from collections import defaultdict
from lxml import etree, objectify
from xml.sax.saxutils import escape

data_elements = ["recordLookups", "recordUpdates", "recordCreates", "recordDeletes"]
interaction_elements = ["actionCalls", "subflows", "screens"]
logic_elements = ["decisions", "assignments", "waits", "loops", "collectionProcessors"]
elements = interaction_elements + logic_elements + data_elements
resources = ["variables", "constants", "formulas", "textTemplates", "stages"]
array_keys = [
    "actionCalls", "apexPluginCalls", "assignments", "choices", "collectionProcessors",
    "constants", "decisions", "dynamicChoiceSets", "formulas", "loops", "orchestratedStages",
    "processMetadataValues", "recordCreates", "recordDeletes", "recordLookups", 
    "recordRollbacks", "recordUpdates", "screens", "stages", "steps", "subflows",
    "textTemplates", "variables", "waits", "dataTypeMappings", "inputParameters",
    "outputParameters", "assignmentItems", "conditions", "mapItems", "sortOptions",
    "rules", "filters", "outputAssignments", "processMetadataValues", "exitActionInputParameters",
    "exitActionOutputParameters", "exitConditions", "stageSteps", "inputAssignments",
    "choiceReferences", "assignees", "entryActionInputParameters", "entryActionOutputParameters",
    "entryConditions", "scheduledPaths", "connectors", "waitEvents"
              ]

def parse(file_name, remove_blank_text=False):
    parser = etree.XMLParser(remove_blank_text=remove_blank_text)
    tree = etree.parse(file_name, parser)
    return tree

def remove_namespace(tree):
    root = tree.getroot()
    for elem in root.getiterator():
        if not hasattr(elem.tag, 'find'): continue  # guard for Comment tags
        i = elem.tag.find('}')
        if i >= 0:
            elem.tag = elem.tag[i+1:]
    objectify.deannotate(root, cleanup_namespaces=True)
    return tree

def xml2dict(root, d={}):
    if root is None: return
    childrens = root.getchildren()
    if len(childrens) == 0:
        if root.tag.startswith("location"): return
        if root.tag in array_keys:
            if root.tag in d:
                d[root.tag].append(root.text)
            else:
                d[root.tag] = [root.text]
        else:
            d[root.tag] = root.text
        return
    t = {}
    for child in childrens:
        xml2dict(child, t)
    if root.tag in array_keys:
        if root.tag in d:
            d[root.tag].append(t)
        else:
            d[root.tag] = [t]
    else:
        d[root.tag] = t

def find_target(json_input, lookup_key):
    if isinstance(json_input, dict):
        for k, v in json_input.items():
            if k == lookup_key:
                yield v
            else:
                yield from find_target(v, lookup_key)
    elif isinstance(json_input, list):
        for item in json_input:
            yield from find_target(item, lookup_key)

def convert_format(dict_input):
    res = {"elements": [], "connectors": [], "resources": [], "properties": []}
    for cpn_type, cpn_val in dict_input.items():
        if cpn_type in elements:
            cpn_vals = cpn_val if isinstance(cpn_val, list) else [cpn_val]
            for element in cpn_vals:
                transformed = {"elementType": cpn_type, "parameters": []}
                for key, val in element.items():
                    if key == "name":
                        transformed["name"] = val
                    else:
                        transformed["parameters"].append({key: val})
                res["elements"].append(transformed)
                targets = find_target(element, "targetReference")
                for t in targets:
                    res["connectors"].append({"source": element["name"], "target": t})
        elif cpn_type in resources:
            cpn_vals = cpn_val if isinstance(cpn_val, list) else [cpn_val]
            for resource in cpn_vals:
                transformed = {"resourceType": cpn_type, "parameters": []}
                for key, val in resource.items():
                    if key == "name":
                        transformed["name"] = val
                    else:
                        transformed["parameters"].append({key: val})
                res["resources"].append(transformed)
        else:
            res["properties"].append({cpn_type: cpn_val})
    return res

def convert_topology(dict_input):
    res = {"elements": [], "connectors": [], "resources": [], "properties": {}}
    for cpn_type, cpn_val in dict_input.items():
        if cpn_type in elements:
            for element in cpn_val:
                transformed = {"elementType": cpn_type, "name": element["name"]}
                res["elements"].append(transformed)
                targets = find_target(element, "targetReference")
#                 for t in targets:
#                     res["connections"].append({"source": element["name"], "target": t})
                connector = find_target(element, "connector")
                for t in connector:
                    res["connectors"].append({"source": element["name"], "target": t["targetReference"]})
                connector = find_target(element, "defaultConnector")
                for t in connector:
                    res["connectors"].append({"source": element["name"], "defaultTarget": t["targetReference"]})
        elif cpn_type in resources:
            for resource in cpn_val:
                transformed = {"resourceType": cpn_type, "resourceBlob": resource}
                res["resources"].append(transformed)
        else:
#             if cpn_type in ["start", "description", "apiVersion"]:
             res["properties"][cpn_type] = cpn_val
    return res

def convert_element_details(dict_input):
    res = {"elements": []}
    for cpn_type, cpn_val in dict_input.items():
        if cpn_type in elements:
            for element in cpn_val:
                transformed = {"elementType": cpn_type, "elementBlob": element}
                res["elements"].append(transformed)
    return res

def compose(topology, details):
    d = {"Flow": {}}
    for e in details["elements"]:
        ele_type = e["elementType"]
        ele_blob = e["elementBlob"]
        ele_blob["locationX"] = "374"
        ele_blob["locationY"] = "288"
        if ele_type in d["Flow"]:
            d["Flow"][ele_type].append(ele_blob)
        else:
            d["Flow"][ele_type] = [ele_blob]
    for r in topology["resources"]:
        res_type = r["resourceType"]
        res_blob = r["resourceBlob"]
        if res_type in d["Flow"]:
            d["Flow"][res_type].append(res_blob)
        else:
            d["Flow"][res_type] = [res_blob]
    for p, v in topology["properties"].items():
        if p == "start":
            v["locationX"] = "50"
            v["locationY"] = "0"
        d["Flow"][p] = v
    return d

def dict2xml(d, level=0):
    ans = ""
    indent = " " * level
    if not d: return ans
    for k, v in d.items():
        if isinstance(v, dict):
            ans += f"{indent}<{k}>\n" + dict2xml(v, level + 2) + f"{indent}</{k}>\n"
        elif isinstance(v, list):
            for v1 in v:
                if isinstance(v1, dict):
                    ans += f"{indent}<{k}>\n" + dict2xml(v1, level + 2) + f"{indent}</{k}>\n"
                else:
                    ev1 = escape(v1) if isinstance(v1, str) else v1
                    ans += f"{indent}<{k}>{ev1}</{k}>\n"
        else:
            ev = escape(v) if isinstance(v, str) else v
            ans += f"{indent}<{k}>{ev}</{k}>\n"
    return ans