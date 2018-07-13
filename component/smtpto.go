package component

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"
)

/*邮件配置部分*/

// mail配置
const (
	Host     = "smtp.163.com"
	Port     = 25
	UserMail = "xxx@163.com"
	PassWord = "xxxx"
	From     = "xxx@163.com"
)

// CommentMailTemp 评论的模板
var CommentMailTemp *template.Template

func init() {
	var err error
	CommentMailTemp, err = template.New("commment_tpl").Parse(MailTemplate)
	if err != nil {
		panic(err)
	}
}

// 邮件的配置模板
const (
	MailTemplate = `<center>
	<table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" id="bodyTable" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;margin: 0;padding: 0;background-color: #ededed; width: 100% !important;">
		<tr >
			<td align="center" valign="top" id="bodyCell" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;margin: 0;padding: 0;border-top: 0;height: 100% !important;width: 100% !important;">
				<table class="setting-max-width" border="0" cellpadding="0" cellspacing="0" width="600px" style="margin:10px 0 10px 0;border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;background-color:#ffffff; width: 600px;">
					<tr>
						<td id="email-header" align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">

							<table border="0" cellpadding="0" cellspacing="0" class="text" width="100%" id="templateBody" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;background-color: transparent
 ;margin-top:0px;margin-bottom:0px;">
								<tr>
									<td align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
										<table border="0" cellpadding="0" cellspacing="0" width="100%" class="templateContainer"
											style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
											<tr>
												<td valign="top" class="bodyContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
													<table border="0" cellpadding="0" cellspacing="0" width="100%" class="mcnTextBlock"
														style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
														<tbody class="mcnTextBlockOuter">
															<tr>
																<td valign="top" class="mcnTextBlockInner" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																	<table align="left" border="0" cellpadding="0" cellspacing="0"
																		width="100%" class="mcnTextContentContainer" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																		<tbody>
																			<tr style="background-color: #539C65!important;">

																				<td valign="top" class="mcnTextContent" style="padding-top:15px; padding-bottom:15px; padding-left:15px; padding-right:15px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;text-align: left;font-size:13px;font-family:&#39;Microsoft Yahei&#39;;color:rgb(51, 51, 51);">
																					<p>
																						<strong>
																							<span style="font-size: 20px; color: rgb(255, 255, 255);">{{.sitename}}</span>
																						</strong>
																					</p>
																					<p>
																						<span style="font-size: 20px; color: rgb(255, 255, 255);">
																							<strong>&nbsp;&nbsp;&nbsp;&nbsp;
																								<span
																									style="color: rgb(255, 255, 255); font-size: 14px;">{{.sitesignature}}</span>
																						<br />
																						</strong>
																						</span>
																					</p>
																				</td>
																			</tr>
																		</tbody>
																	</table>

																</td>
															</tr>
														</tbody>
													</table>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td id="email-body" align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">

							<table border="0" cellpadding="0" cellspacing="0" class="text" width="100%" id="templateBody" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;background-color: transparent
 ;margin-top:0px;margin-bottom:0px;">
								<tr>
									<td align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
										<table border="0" cellpadding="0" cellspacing="0" width="100%" class="templateContainer"
											style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
											<tr>
												<td valign="top" class="bodyContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
													<table border="0" cellpadding="0" cellspacing="0" width="100%" class="mcnTextBlock"
														style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
														<tbody class="mcnTextBlockOuter">
															<tr>
																<td valign="top" class="mcnTextBlockInner" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																	<table align="left" border="0" cellpadding="0" cellspacing="0"
																		width="100%" class="mcnTextContentContainer" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																		<tbody>
																			<tr>

																				<td valign="top" class="mcnTextContent" style="padding-top:15px; padding-bottom:15px; padding-left:15px; padding-right:15px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;text-align: left;font-size:13px;font-family:&#39;Microsoft Yahei&#39;;color:rgb(51, 51, 51);">
																					<p style="text-align: center;">
																						<span style="font-size: 16px;"></span>
																						<span style="font-size: 18px; color: rgb(103, 147, 116);">
																							<strong>有新的评论,请注意查看。</strong>
																						</span>
																					</p>
																					<p>
																						<span style="color: rgb(127, 127, 127); font-size: 15px;">
																							<br />
																						</span>
																					</p>
																					<p>
																						<span style="color:#a5a5a5">
																							<span style="font-size: 15px;">{{.username}},您好！</span>
																						</span>
																					</p>
																					<p>
																						<span style="color:#a5a5a5">
																							<span style="font-size: 15px;">
																								<br />
																							</span>
																						</span>
																					</p>
																					<ul class="custom_dash list-paddingleft-1">
																						<li class="list-dash list-dash-paddingleft">
																							<p>
																								<span style="font-size: 15px; color: rgb(165, 165, 165);">你在{{.sitename}}发布的{{.arttitle}}
																									<strong>{{.arttitle}} </strong>有新的评论,需要审核。</span>
																							</p>
																						</li>
																						<li class="list-dash list-dash-paddingleft">
																							<p>
																								<span style="color:#a5a5a5">
																									<span style="font-size: 15px;">以下是详细内容：</span>
																								</span>
																							</p>
																						</li>
																						<li class="list-dash list-dash-paddingleft">
																							<table style="border-collapse:collapse;"
																								interlaced="enabled" align="center"
																								data-sort="sortDisabled">
																								<tbody>
																									<tr class="ue-table-interlace-color-single firstRow">
																										<td width="407" valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">文章地址</span>
																										</td>
																										<td width="407"
																											valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">{{.arturl}}
																												<br
																												/>
																											</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-double">
																										<td width="407" valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">评论人</span>
																										</td>
																										<td width="407"
																											valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">{{.author}}
																												<br
																												/>
																											</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-single">
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">邮箱
																												<br />
																											</span>
																										</td>
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">{{.mail}}
																												<br
																												/>
																											</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-double">
																										<td width="407" valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">域名</span>
																										</td>
																										<td width="407"
																											valign="top"
																											style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																											<span style="color: rgb(127, 127, 127);">{{.url}}</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-double">
																											<td width="407" valign="top"
																												style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																												<span style="color: rgb(127, 127, 127);">Ip</span>
																											</td>
																											<td width="407"
																												valign="top"
																												style="border-width: 1px; border-right-style: solid; border-bottom-style: solid; border-color: rgb(221, 221, 221); border-image: initial; word-break: break-all;">
																												<span style="color: rgb(127, 127, 127);">{{.authorip}}</span>
																											</td>
																										</tr>
																									<tr class="ue-table-interlace-color-single">
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">评论内容</span>
																										</td>
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">{{.content}}<br/>
																											</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-double">
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">时间</span>
																										</td>
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">{{.nowdate}}</span>
																										</td>
																									</tr>
																									<tr class="ue-table-interlace-color-single">
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);">审核操作</span>
																										</td>
																										<td style="border: 1px solid rgb(221, 221, 221); word-break: break-all;"
																											width="407" valign="top">
																											<span style="color: rgb(127, 127, 127);"><a style="color:#539C65;" href="{{.passurl}}">通过</a>/<a style="color:#539C65;" href="{{.delurl}}">删除</a></span>
																										</td>
																									</tr>
																								</tbody>
																							</table>
																						</li>
																						<li class="list-dash list-dash-paddingleft">
																							<p>
																								<span style="color:#a5a5a5;">你还有 <b style="color:black;" >{{.num}}</b> 个评论没有审核,详情查看:点击 <a style="color:#539C65;" href="{{.untreatedlist}}">此处</a></span>
																							</p>
																						</li>
																					</ul>
																					<p></p>
																				</td>
																			</tr>
																		</tbody>
																	</table>
																</td>
															</tr>
														</tbody>
													</table>
												</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td id="email-footer" align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
							<table border="0" cellpadding="0" cellspacing="0" width="100%" id="templateFooter" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;background-color:     #fff ;margin-top:0px;margin-bottom:0px;">
								<tr>
									<td valign="top" class="footerContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
										<table border="0" cellpadding="0" cellspacing="0" width="100%" class="mcnButtonBlock"
											style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
											<tbody class="mcnButtonBlockOuter">
												<tr>
													<td style="padding-top:15px; padding-bottom:15px; padding-left:15px; padding-right:15px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;text-align:center"
														valign="top" align="center" class="mcnButtonBlockInner">
														<a target="_blank" href="http://www.yulibaozi.com"
															style="word-wrap: break-word;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;padding: 5px 0px;background-color: #7abce7;margin: 1px 0;font-size: 14px;line-height: 1.5;font-weight: normal;text-align: center;vertical-align: middle;cursor: pointer;background-image: none;display:inline-block;border-radius: 3px;text-decoration:none;width:150px;min-height:23px;line-height:23px;background:#e6e6e6;color:#679374"
															class="btn-sm btn-info">yulibaozi.com</a>
													</td>
												</tr>
											</tbody>
										</table>
									</td>
								</tr>
							</table>
							<table border="0" cellpadding="0" cellspacing="0" width="100%" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;background-color:     #fff ;margin-top:0px;margin-bottom:0px;">
								<tr>
									<td valign="top" class="preheaderContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
										<table border="0" cellpadding="0" cellspacing="0" width="100%" class="mcnDividerBlock"
											style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
											<tbody class="mcnDividerBlockOuter">
												<tr>
													<td class="mcnDividerBlockInner" style="padding-top:15px; padding-bottom:15px; padding-left:15px; padding-right:15px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
														<table class="mcnDividerContent" border="0" cellpadding="0" cellspacing="0"
															width="100%" style="border-top-width: 1px;border-top-style: dashed
						;border-top-color: #d4d4d4;border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
															<tbody>
																<tr>
																	<td style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																		<span></span>
																	</td>
																</tr>
															</tbody>
														</table>
													</td>
												</tr>
											</tbody>
										</table>
									</td>
								</tr>
							</table>
							<table border="0" cellpadding="0" cellspacing="0" width="100%" id="templateColumns" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;">
								<tbody>
									<tr class="layer-wrap" style="background:transparent
">
										<td align="left" valign="top" class="columnsContainer two-column" width="60%" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
											<table border="0" cellpadding="0" cellspacing="0" width="100%" class="templateColumn"
												style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
												<tbody>
													<tr>
														<td valign="top" class="leftColumnContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
															<table border="0" cellpadding="0" cellspacing="0" width="100%"
																class="mcnTextBlock" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																<tbody class="mcnTextBlockOuter">
																	<tr>
																		<td valign="top" class="mcnTextBlockInner" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																			<table align="left" border="0" cellpadding="0"
																				cellspacing="0" width="100%" class="mcnTextContentContainer"
																				style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																				<tbody>
																					<tr>

																						<td valign="top" class="mcnTextContent content-wrap" id="content-wrap" style=" mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;color: rgb(51, 51, 51);text-align: left ">
																							<table border="0" cellpadding="0" cellspacing="0" class="text" width="100%" id="templateBody" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;border-top: 0;border-bottom: 0;background-color: transparent;margin-top:0px;margin-bottom:0px;">
																								<tr>
																									<td align="center" valign="top" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																										<table border="0"
																											cellpadding="0"
																											cellspacing="0"
																											width="100%"
																											class="templateContainer"
																											style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																											<tr>
																												<td valign="top" class="bodyContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																													<table
																														border="0"
																														cellpadding="0"
																														cellspacing="0"
																														width="100%"
																														class="mcnTextBlock"
																														style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																														<tbody class="mcnTextBlockOuter">
																															<tr>
																																<td valign="top" class="mcnTextBlockInner" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																																	<table
																																		align="left"
																																		border="0"
																																		cellpadding="0"
																																		cellspacing="0"
																																		width="100%"
																																		class="mcnTextContentContainer"
																																		style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
																																		<tbody>
																																			<tr>

																																				<td valign="top" class="mcnTextContent" style="padding-top:15px; padding-bottom:15px; padding-left:15px; padding-right:15px;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;text-align: left;font-size:13px;font-family:&#39;Microsoft Yahei&#39;;color:rgb(51, 51, 51);">
																																					<p>
																																						<span
																																							style="color:#7f7f7f">{{.sitesignature}}</span>
																																					</p>
																																					<p>
																																						<span
																																							style="color:#7f7f7f">
																																							<br
																																							/>
																																							</span>
																																					</p>
																																					<p>
																																						<span
																																							style="color: rgb(127, 127, 127);">E-Mail:{{.useremail}}</span>
																																					</p>
																																					<p>
																																						<span
																																							style="color: rgb(127, 127, 127);">网址：www.yulibaozi.com</span>
																																					</p>
																																					<p>
																																						<span
																																							style="color: rgb(127, 127, 127);">地址：中国</span>
																																					</p>
																																				</td>
																																			</tr>
																																		</tbody>
																										</table>

																										</td>
																										</tr>
																										</tbody>
																							</table>
																							</td>
																							</tr>
																			</table>
																			</td>
																			</tr>
															</table>
															</td>
															</tr>
															</tbody>
											</table>

											</td>
											</tr>
											</tbody>
							</table>
							</td>
							</tr>
							</tbody>
				</table>
				</td>
				<td align="left" valign="top" class="columnsContainer two-column" width="40%" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
					<table border="0" cellpadding="0" cellspacing="0" width="100%" class="templateColumn" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
						<tbody>
							<tr>
								<td valign="top" class="leftColumnContainer" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
									<table border="0" cellpadding="0" cellspacing="0" width="100%" class="mcnTextBlock" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
										<tbody class="mcnTextBlockOuter">
											<tr>
												<td valign="top" class="mcnTextBlockInner" style="mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
													<table align="left" border="0" cellpadding="0" cellspacing="0" width="100%"
														class="mcnTextContentContainer" style="border-collapse: collapse;mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;">
														<tbody>
															<tr>

																<td valign="top" class="mcnTextContent content-wrap" id="content-wrap" style=" mso-table-lspace: 0pt;mso-table-rspace: 0pt;-ms-text-size-adjust: 100%;-webkit-text-size-adjust: 100%;color: rgb(51, 51, 51);text-align: left ">
																</td>
															</tr>
														</tbody>
													</table>

												</td>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
				</td>
				</tr>
				</tbody>
	</table>
	</td>
	</tr>
	</table>

	</td>
	</tr>
	</table>
</center>
	`
)

/*
站点名字 {{.sitename}}
	站点签名 {{.sitesignature}}
	用户名 {{.username}}
	用户邮箱 {{.useremail}}
	文章名 {{.arttitle}}
	文章地址 {{.arturl}}
	评论人名称 {{.author}},邮箱 {{.mail}},域名 {{.url}},IP {{.authorip}},内容 {{.content}},两个操作URL {{.passurl}}/{{.delurl}}
	评论时间 {{.nowdate}}


	还有几个没有发 {{.num}},以及未处理列表 {{.untreatedlist}}
*/

/*发送部分*/

var (
	sendCh chan *Email
)

func init() {
	go func() {
		sendCh = make(chan *Email, 1000)
		for {
			select {
			case m, ok := <-sendCh:
				if !ok {
					return
				}
				if err := m.Send(); err != nil {
					fmt.Println("发送email出错:" + err.Error())
				}
			}
		}
	}()
}

// Param 变量
type Param struct {
	Address string `json:"address"` //邮件地址
	Title   string `json:"title"`   //邮件标题
	Content string `json:"content"` //邮件内容
}

var confStr = getConf(UserMail, PassWord, Host, From, Port)

// getConf 获取配置
func getConf(user, pw, host, from string, port int) string {
	return fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d,"from":"%s"}`, user, pw, host, port, from)
}

// SendMail 发送邮件
func SendMail(p *Param) error {
	email := NewEMail(confStr)
	email.Subject = p.Title
	email.HTML = p.Content
	email.To = []string{p.Address}
	select {
	case sendCh <- email:
		return nil
	case <-time.After(time.Second * 10):
		return errors.New("发送失败")
	}
}

// ExecTemp 和HTML代码合并生成模板
func ExecTemp(data map[string]interface{}) (msg string, err error) {
	conent := new(bytes.Buffer)
	err = CommentMailTemp.Execute(conent, data)
	if err != nil {
		return "", err
	}
	return conent.String(), nil
}
